package pkg

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const testAnnotation = "alb.ingress.kubernetes.io/test"

// Validate checks the ingress annotation test is exists
func Validate(ing metav1.Object) error {
	_, found := ing.GetAnnotations()[testAnnotation]
	if !found {
		return fmt.Errorf("deny")
	}
	return nil
}
