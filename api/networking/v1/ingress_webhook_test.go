package v1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	netv1 "k8s.io/api/networking/v1"
)

const testAnnotation = "alb.ingress.kubernetes.io/test"

var _ = Describe("valid cases for Ingress validator", func() {
	It("should allow creating Ingress with annotated test", func() {
		ing := &netv1.Ingress{}
		ing.Name = "allow-creating-namespaced"
		ing.Namespace = "default"
		ing.Annotations = map[string]string{testAnnotation: "hogefuga"}
		ing.Spec.DefaultBackend = &netv1.IngressBackend{Service: &netv1.IngressServiceBackend{Name: "test", Port: netv1.ServiceBackendPort{
			Number: 8080,
		}}}
		err := k8sClient.Create(ctx, ing)
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("invalid cases for Ingress", func() {
	It("should deny creating Ingress without annotated test", func() {
		ing := &netv1.Ingress{}
		ing.Name = "deny-creating-not-namespaced"
		ing.Namespace = "default"
		// ing.Annotations = map[string]string{groupNameAnnotation: "test"}
		ing.Spec.DefaultBackend = &netv1.IngressBackend{Service: &netv1.IngressServiceBackend{Name: "test", Port: netv1.ServiceBackendPort{
			Number: 8080,
		}}}
		err := k8sClient.Create(ctx, ing)
		Expect(err).To(HaveOccurred())
	})
})
