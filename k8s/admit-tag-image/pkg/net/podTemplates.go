package net

import (
	"fmt"
	"github.com/dual-lab/admit-webook-boilerplate/pkg/webhook"
	"github.com/dual-lab/admit-webook-boilerplate/pkg/webhook/net"
	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"net/http"
)

type podTemplatesFunc webhook.AdmitV1Func

func podTemplatesFuncGet() podTemplatesFunc {
	return internalPodTemplatesAdmit
}

func (r podTemplatesFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	net.Serve(w, req, webhook.WrapToAdminV1(webhook.AdmitV1Func(r)))
}

func internalPodTemplatesAdmit(review v1.AdmissionReview) *v1.AdmissionResponse {
	klog.V(2).Info("admit podTemplates")
	groupVersionResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "podtemplates"}
	if review.Request.Resource != groupVersionResource {
		err := admitErroNew(fmt.Errorf("expexted resource to be %s", groupVersionResource))
		klog.Error(err)
		return err.toAdmitResponse()
	}

	raw := review.Request.Object.Raw
	podTemplate := corev1.PodTemplate{}

	des := webhook.Codecs.UniversalDeserializer()
	if _, _, err := des.Decode(raw, nil, &podTemplate); err != nil {
		klog.Error(err)
		return admitErroNew(err).toAdmitResponse()
	}
	reviewResponse := v1.AdmissionResponse{}
	reviewResponse.Allowed = true

	// Ignore if present ignore flag on podTemplate labels
	if hasIgnoreNotLatestImageLabels(podTemplate.Template.Labels) {
		reviewResponse.Result = &metav1.Status{Message: "Ignore not latest image policy"}
	} else {
		msg, ok := haveSomeImageLatestTag(podTemplate.Template.Spec.Containers)
		reviewResponse.Allowed = !ok
		reviewResponse.Result = &metav1.Status{Message: msg}
	}
	return &reviewResponse
}
