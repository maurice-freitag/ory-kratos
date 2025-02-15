package template_test

import (
	"context"
	"testing"

	"github.com/ory/kratos/courier"
	"github.com/ory/kratos/courier/template/testhelpers"

	"github.com/ory/kratos/courier/template"
	"github.com/ory/kratos/internal"
)

func TestVerifyValid(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	t.Run("test=with courier templates directory", func(t *testing.T) {
		_, reg := internal.NewFastRegistryWithMocks(t)
		tpl := template.NewVerificationValid(reg, &template.VerificationValidModel{})

		testhelpers.TestRendered(t, ctx, tpl)
	})

	t.Run("test=with remote resources", func(t *testing.T) {
		testhelpers.TestRemoteTemplates(t, "courier/builtin/templates/verification/valid", courier.TypeVerificationValid)
	})
}
