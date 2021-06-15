package verification

import (
	"net/http"
	"time"

	"github.com/ory/kratos/schema"
	"github.com/ory/kratos/ui/node"
	"github.com/ory/x/sqlcon"

	"github.com/ory/herodot"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/config"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/selfservice/flow"
	"github.com/ory/kratos/x"
)

const (
	RouteInitBrowserFlow = "/self-service/verification/browser"
	RouteInitAPIFlow     = "/self-service/verification/api"
	RouteGetFlow         = "/self-service/verification/flows"

	RouteSubmitFlow = "/self-service/verification"
)

type (
	HandlerProvider interface {
		VerificationHandler() *Handler
	}
	handlerDependencies interface {
		errorx.ManagementProvider
		identity.ManagementProvider
		identity.PrivilegedPoolProvider
		config.Provider

		x.CSRFTokenGeneratorProvider
		x.WriterProvider
		x.CSRFProvider

		FlowPersistenceProvider
		ErrorHandlerProvider
		StrategyProvider
	}
	Handler struct {
		d handlerDependencies
	}
)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

func (h *Handler) RegisterPublicRoutes(public *x.RouterPublic) {
	h.d.CSRFHandler().IgnorePath(RouteInitAPIFlow)
	h.d.CSRFHandler().IgnorePath(RouteSubmitFlow)

	public.GET(RouteInitBrowserFlow, h.initBrowserFlow)
	public.GET(RouteInitAPIFlow, h.initAPIFlow)
	public.GET(RouteGetFlow, h.fetch)

	public.POST(RouteSubmitFlow, h.submitFlow)
	public.GET(RouteSubmitFlow, h.submitFlow)
}

func (h *Handler) RegisterAdminRoutes(admin *x.RouterAdmin) {
	admin.GET(RouteGetFlow, h.fetch)
}

// swagger:route GET /self-service/verification/api public initializeSelfServiceVerificationWithoutBrowser
//
// Initialize Verification Flow for APIs, Services, Apps, ...
//
// :::info
//
// This endpoint is EXPERIMENTAL and subject to potential breaking changes in the future.
//
// :::
//
// This endpoint initiates a verification flow for API clients such as mobile devices, smart TVs, and so on.
//
// To fetch an existing verification flow call `/self-service/verification/flows?flow=<flow_id>`.
//
// :::warning
//
// You MUST NOT use this endpoint in client-side (Single Page Apps, ReactJS, AngularJS) nor server-side (Java Server
// Pages, NodeJS, PHP, Golang, ...) browser applications. Using this endpoint in these applications will make
// you vulnerable to a variety of CSRF attacks.
//
// This endpoint MUST ONLY be used in scenarios such as native mobile apps (React Native, Objective C, Swift, Java, ...).
//
// :::
//
// More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       500: jsonError
//       400: jsonError
func (h *Handler) initAPIFlow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !h.d.Config(r.Context()).SelfServiceFlowVerificationEnabled() {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Verification is not allowed because it was disabled.")))
		return
	}

	req, err := NewFlow(h.d.Config(r.Context()), h.d.Config(r.Context()).SelfServiceFlowVerificationRequestLifespan(), h.d.GenerateCSRFToken(r), r, h.d.VerificationStrategies(r.Context()), flow.TypeAPI)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if err := h.d.VerificationFlowPersister().CreateVerificationFlow(r.Context(), req); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, req)
}

// swagger:route GET /self-service/verification/browser public initializeSelfServiceVerificationForBrowsers
//
// Initialize Browser-Based Verification Flow
//
// :::info
//
// This endpoint is EXPERIMENTAL and subject to potential breaking changes in the future.
//
// :::
//
// This endpoint initializes a browser-based account verification flow. Once initialized, the browser will be redirected to
// `selfservice.flows.verification.ui_url` with the flow ID set as the query parameter `?flow=`.
//
// If this endpoint is called via an AJAX request, the response contains the recovery flow without any redirects.
//
// This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).
//
// More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       302: emptyResponse
//       500: jsonError
func (h *Handler) initBrowserFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !h.d.Config(r.Context()).SelfServiceFlowVerificationEnabled() {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Verification is not allowed because it was disabled.")))
		return
	}

	req, err := NewFlow(h.d.Config(r.Context()), h.d.Config(r.Context()).SelfServiceFlowVerificationRequestLifespan(), h.d.GenerateCSRFToken(r), r, h.d.VerificationStrategies(r.Context()), flow.TypeBrowser)
	if err != nil {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	if err := h.d.VerificationFlowPersister().CreateVerificationFlow(r.Context(), req); err != nil {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	redirTo := req.AppendTo(h.d.Config(r.Context()).SelfServiceFlowVerificationUI()).String()
	x.AcceptToRedirectOrJson(w, r, h.d.Writer(), req, redirTo)
}

// nolint:deadcode,unused
// swagger:parameters getSelfServiceVerificationFlow
type getSelfServiceVerificationFlowParameters struct {
	// The Flow ID
	//
	// The value for this parameter comes from `request` URL Query parameter sent to your
	// application (e.g. `/verification?flow=abcde`).
	//
	// required: true
	// in: query
	FlowID string `json:"id"`
}

// swagger:route GET /self-service/verification/flows public admin getSelfServiceVerificationFlow
//
// Get Verification Flow
//
// :::info
//
// This endpoint is EXPERIMENTAL and subject to potential breaking changes in the future.
//
// :::
//
// This endpoint returns a verification flow's context with, for example, error details and other information.
//
// More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       403: jsonError
//       404: jsonError
//       500: jsonError
func (h *Handler) fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !h.d.Config(r.Context()).SelfServiceFlowVerificationEnabled() {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Verification is not allowed because it was disabled.")))
		return
	}

	rid := x.ParseUUID(r.URL.Query().Get("id"))
	req, err := h.d.VerificationFlowPersister().GetVerificationFlow(r.Context(), rid)
	if err != nil {
		return err
	}

	if mustVerify && !nosurf.VerifyToken(h.d.GenerateCSRFToken(r), ar.CSRFToken) {
		return errors.WithStack(x.ErrInvalidCSRFToken.WithDebugf("Expected %s but got %s", h.d.GenerateCSRFToken(r), ar.CSRFToken))
	}

	h.d.Writer().Write(w, r, ar)
	return nil
}

// nolint:deadcode,unused
// swagger:parameters completeSelfServiceBrowserVerificationFlow
type completeSelfServiceBrowserVerificationFlowParameters struct {
	// Request is the Request ID
	//
	// The value for this parameter comes from `request` URL Query parameter sent to your
	// application (e.g. `/verify?request=abcde`).
	//
	// required: true
	// in: query
	Request string `json:"request"`

	// What to verify
	//
	// Currently only "email" is supported.
	//
	// required: true
	// in: path
	Via string `json:"via"`
}

// swagger:route POST /self-service/browser/flows/verification/{via}/complete public completeSelfServiceBrowserVerificationFlow
//
// Complete the Browser-Based Verification Flows
//
// This endpoint completes a browser-based verification flow. This is usually achieved by POSTing data to this
// endpoint.
//
// If the provided data is valid against the Identity's Traits JSON Schema, the data will be updated and
// the browser redirected to `url.settings_ui` for further steps.
//
// > This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.
//
// More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Consumes:
//     - application/json
//     - application/x-www-form-urlencoded
//
//     Schemes: http, https
//
//     Responses:
//       302: emptyResponse
//       500: genericError
func (h *Handler) complete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if _, err := h.toVia(ps); err != nil {
		h.handleError(w, r, nil, err)
		return
	}

	if req.ExpiresAt.Before(time.Now().UTC()) {
		if req.Type == flow.TypeBrowser {
			h.d.Writer().WriteError(w, r, errors.WithStack(x.ErrGone.
				WithReason("The verification flow has expired. Redirect the user to the verification flow init endpoint to initialize a new verification flow.").
				WithDetail("redirect_to", urlx.AppendPaths(h.d.Config(r.Context()).SelfPublicURL(r), RouteInitBrowserFlow).String())))
			return
		}
		h.d.Writer().WriteError(w, r, errors.WithStack(x.ErrGone.
			WithReason("The verification flow has expired. Call the verification flow init API endpoint to initialize a new verification flow.").
			WithDetail("api", urlx.AppendPaths(h.d.Config(r.Context()).SelfPublicURL(r), RouteInitAPIFlow).String())))
		return
	}

	http.Redirect(w, r, h.c.SelfServiceFlowVerificationReturnTo().String(), http.StatusFound)
}

// nolint:deadcode,unused
// swagger:parameters selfServiceBrowserVerify
type selfServiceBrowserVerifyParameters struct {
	// required: true
	// in: path
	Code string `json:"code"`

	// What to verify
	//
	// Currently only "email" is supported.
	//
	// required: true
	// in: path
	Via string `json:"via"`
}

// swagger:route GET /self-service/browser/flows/verification/{via}/confirm/{code} public selfServiceBrowserVerify
//
// Complete the Browser-Based Verification Flows
//
// This endpoint completes a browser-based verification flow.
//
// > This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.
//
// More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Consumes:
//     - application/json
//     - application/x-www-form-urlencoded
//
//     Schemes: http, https
//
//     Responses:
//       302: emptyResponse
//       500: genericError
func (h *Handler) verify(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	via, err := h.toVia(ps)
	if err != nil {
		h.handleError(w, r, nil, err)
		return
	}

	if err := h.d.PrivilegedIdentityPool().VerifyAddress(r.Context(), ps.ByName("code")); err != nil {
		if errors.Is(err, sqlcon.ErrNoRows) {
			a := NewRequest(
				h.c.SelfServiceFlowSettingsRequestLifespan(), r, via,
				urlx.AppendPaths(h.c.SelfPublicURL(), strings.ReplaceAll(PublicVerificationCompletePath, ":via", string(via))), h.d.GenerateCSRFToken,
			)

			a.Messages.Add(text.NewErrorValidationVerificationTokenInvalidOrAlreadyUsed())
			if err := h.d.VerificationPersister().CreateVerificationRequest(r.Context(), a); err != nil {
				h.handleError(w, r, nil, err)
				return
			}

			http.Redirect(w, r,
				urlx.CopyWithQuery(h.c.SelfServiceFlowVerificationUI(), url.Values{"request": {a.ID.String()}}).String(),
				http.StatusFound,
			)
			return
		}

		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	http.Redirect(w, r, h.c.SelfServiceFlowVerificationReturnTo().String(), http.StatusFound)
}

// handleError is a convenience function for handling all types of errors that may occur (e.g. validation error).
func (h *Handler) handleError(w http.ResponseWriter, r *http.Request, rr *Request, err error) {
	if rr != nil {
		rr.Form.Reset()
		rr.Form.SetCSRF(h.d.GenerateCSRFToken(r))
	}

	h.d.VerificationRequestErrorHandler().HandleVerificationError(w, r, rr, err)
}

func (h *Handler) toVia(ps httprouter.Params) (identity.VerifiableAddressType, error) {
	v := ps.ByName("via")
	switch identity.VerifiableAddressType(v) {
	case identity.VerifiableAddressTypeEmail:
		return identity.VerifiableAddressTypeEmail, nil
	}
	return "", errors.WithStack(herodot.ErrBadRequest.WithReasonf("Verification only works for email but got: %s", v))
}

// nolint:deadcode,unused
// swagger:parameters submitSelfServiceVerificationFlow
type submitSelfServiceVerificationFlow struct {
	// The Registration Flow ID
	//
	// The value for this parameter comes from `flow` URL Query parameter sent to your
	// application (e.g. `/registration?flow=abcde`).
	//
	// required: true
	// in: query
	Flow string `json:"flow"`

	// in: body
	Body submitSelfServiceRecoveryFlowBody
}

// swagger:model submitSelfServiceRecoveryFlow
// nolint:deadcode,unused
type submitSelfServiceRecoveryFlowBody struct{}

// swagger:route POST /self-service/verification/flows public submitSelfServiceVerificationFlow
//
// Complete Verification Flow
//
// Use this endpoint to complete a verification flow. This endpoint
// behaves differently for API and browser flows and has several states:
//
// - `choose_method` expects `flow` (in the URL query) and `email` (in the body) to be sent
//   and works with API- and Browser-initiated flows.
//	 - For API clients and Browser clients with HTTP Header `Accept: application/json` it either returns a HTTP 200 OK when the form is valid and HTTP 400 OK when the form is invalid
//     and a HTTP 302 Found redirect with a fresh verification flow if the flow was otherwise invalid (e.g. expired).
//	 - For Browser clients without HTTP Header `Accept` or with `Accept: text/*` it returns a HTTP 302 Found redirect to the Verification UI URL with the Verification Flow ID appended.
// - `sent_email` is the success state after `choose_method` when using the `link` method and allows the user to request another verification email. It
//   works for both API and Browser-initiated flows and returns the same responses as the flow in `choose_method` state.
// - `passed_challenge` expects a `token` to be sent in the URL query and given the nature of the flow ("sending a verification link")
//   does not have any API capabilities. The server responds with a HTTP 302 Found redirect either to the Settings UI URL
//   (if the link was valid) and instructs the user to update their password, or a redirect to the Verification UI URL with
//   a new Verification Flow ID which contains an error message that the verification link was invalid.
//
// More information can be found at [Ory Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Consumes:
//     - application/json
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       400: verificationFlow
//       302: emptyResponse
//       500: jsonError
func (h *Handler) submitFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rid, err := flow.GetFlowID(r)
	if err != nil {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, nil, node.DefaultGroup, err)
		return
	}

	f, err := h.d.VerificationFlowPersister().GetVerificationFlow(r.Context(), rid)
	if errors.Is(err, sqlcon.ErrNoRows) {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, nil, node.DefaultGroup, errors.WithStack(herodot.ErrNotFound.WithReasonf("The verification request could not be found. Please restart the flow.")))
		return
	} else if err != nil {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, nil, node.DefaultGroup, err)
		return
	}

	if err := f.Valid(); err != nil {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, f, node.DefaultGroup, err)
		return
	}

	var g node.Group
	var found bool
	for _, ss := range h.d.AllVerificationStrategies() {
		err := ss.Verify(w, r, f)
		if errors.Is(err, flow.ErrStrategyNotResponsible) {
			continue
		} else if errors.Is(err, flow.ErrCompletedByStrategy) {
			return
		} else if err != nil {
			h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, f, ss.VerificationNodeGroup(), err)
			return
		}

		found = true
		g = ss.VerificationNodeGroup()
		break
	}

	if !found {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, f, node.DefaultGroup, errors.WithStack(schema.NewNoVerificationStrategyResponsible()))
		return
	}

	if f.Type == flow.TypeBrowser && !x.IsJSONRequest(r) {
		http.Redirect(w, r, f.AppendTo(h.d.Config(r.Context()).SelfServiceFlowVerificationUI()).String(), http.StatusFound)
		return
	}

	updatedFlow, err := h.d.VerificationFlowPersister().GetVerificationFlow(r.Context(), f.ID)
	if err != nil {
		h.d.VerificationFlowErrorHandler().WriteFlowError(w, r, f, g, err)
		return
	}

	h.d.Writer().Write(w, r, updatedFlow)
}
