package v1

import (
	"context"
	"net/http"

	"github.com/naari3/ingress-validator-test/pkg"
	netv1 "k8s.io/api/networking/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-networking-v1-ingress,mutating=false,failurePolicy=fail,groups="networking.k8s.io",resources=ingresses,verbs=create;update,versions=v1,name=io.naari3.net,sideEffects=None,admissionReviewVersions={v1}

// IngressValidator validates Ingresses
type IngressValidator struct {
	Decoder *admission.Decoder
}

var handleLog = ctrl.Log.WithName("handle")

func (v IngressValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	ing := &netv1.Ingress{}
	handleLog.Info("decoding to networking.k8s.io/Ingress")
	if err := v.Decoder.Decode(req, ing); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := pkg.Validate(ing); err != nil {
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}
