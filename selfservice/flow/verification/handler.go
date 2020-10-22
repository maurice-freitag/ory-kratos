package verification

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/selfservice/flow"
	"github.com/ory/kratos/x"
)

const (
	RouteInitBrowserFlow = "/self-service/verification/browser"
	RouteInitAPIFlow     = "/self-service/verification/api"
	RouteGetFlow         = "/self-service/verification/flows"
)

type (
	HandlerProvider interface {
		VerificationHandler() *Handler
	}
	handlerDependencies interface {
		errorx.ManagementProvider
		identity.ManagementProvider
		identity.PrivilegedPoolProvider
		x.CSRFTokenGeneratorProvider
		x.WriterProvider

		FlowPersistenceProvider
		ErrorHandlerProvider
		StrategyProvider
		x.CSRFProvider
	}
	Handler struct {
		d handlerDependencies
		c configuration.Provider
	}
)

func NewHandler(d handlerDependencies, c configuration.Provider) *Handler {
	return &Handler{c: c, d: d}
}

func (h *Handler) RegisterPublicRoutes(public *x.RouterPublic) {
	h.d.CSRFHandler().ExemptPath(RouteInitAPIFlow)

	public.GET(RouteInitBrowserFlow, h.initBrowserFlow)
	public.GET(RouteInitAPIFlow, h.initAPIFlow)
	public.GET(RouteGetFlow, h.fetch)
}

func (h *Handler) RegisterAdminRoutes(admin *x.RouterAdmin) {
	admin.GET(RouteGetFlow, h.fetch)
}

// swagger:route GET /self-service/verification/api public initializeSelfServiceVerificationViaAPIFlow
//
// Initialize Verification Flow for API Clients
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
// More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       500: genericError
//       400: genericError
func (h *Handler) initAPIFlow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, err := NewFlow(h.c.SelfServiceFlowVerificationRequestLifespan(), h.d.GenerateCSRFToken(r), r, h.d.VerificationStrategies(), flow.TypeAPI)
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

// swagger:route GET /self-service/verification/browser public initializeSelfServiceVerificationViaBrowserFlow
//
// Initialize Browser-Based Verification Flow
//
// This endpoint initializes a browser-based account verification flow. Once initialized, the browser will be redirected to
// `selfservice.flows.verification.ui_url` with the flow ID set as the query parameter `?flow=`.
//
// This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...).
//
// More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Schemes: http, https
//
//     Responses:
//       302: emptyResponse
//       500: genericError
func (h *Handler) initBrowserFlow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req, err := NewFlow(h.c.SelfServiceFlowVerificationRequestLifespan(), h.d.GenerateCSRFToken(r), r, h.d.VerificationStrategies(), flow.TypeBrowser)
	if err != nil {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	if err := h.d.VerificationFlowPersister().CreateVerificationFlow(r.Context(), req); err != nil {
		h.d.SelfServiceErrorManager().Forward(r.Context(), w, r, err)
		return
	}

	http.Redirect(w, r, req.AppendTo(h.c.SelfServiceFlowVerificationUI()).String(), http.StatusFound)
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
// This endpoint returns a verification flow's context with, for example, error details and other information.
//
// More information can be found at [ORY Kratos Email and Phone Verification Documentation](https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation).
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: verificationFlow
//       403: genericError
//       404: genericError
//       500: genericError
func (h *Handler) fetch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
				WithDetail("redirect_to", urlx.AppendPaths(h.c.SelfPublicURL(), RouteInitBrowserFlow).String())))
			return
		}
		h.d.Writer().WriteError(w, r, errors.WithStack(x.ErrGone.
			WithReason("The verification flow has expired. Call the verification flow init API endpoint to initialize a new verification flow.").
			WithDetail("api", urlx.AppendPaths(h.c.SelfPublicURL(), RouteInitAPIFlow).String())))
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
