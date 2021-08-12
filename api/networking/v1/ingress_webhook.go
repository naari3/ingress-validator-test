package v1

import (
	"context"
	"net/http"

	"github.com/naari3/ingress-validator-test/pkg"
	netv1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-networking-v1-ingress,mutating=false,failurePolicy=fail,groups="networking.k8s.io",resources=ingresses,verbs=create;update,versions=v1,name=vingress.kb.io,sideEffects=none,admissionReviewVersions={v1}

// IngressValidator validates Ingresses
type IngressValidator struct {
	decoder *admission.Decoder
}

func (v IngressValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	ing := &netv1.Ingress{}
	if err := v.decoder.Decode(req, ing); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := pkg.ValidateGroupName(ing); err != nil {
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}
