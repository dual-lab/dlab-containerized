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

type podAdmitFunc webhook.AdmitV1Func

func podAdmitFuncGet() podAdmitFunc {
	return internalPodAdmit
}

func (r podAdmitFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	net.Serve(w, req, webhook.WrapToAdminV1(webhook.AdmitV1Func(r)))
}

func internalPodAdmit(review v1.AdmissionReview) *v1.AdmissionResponse {
	klog.V(2).Info("admit pod")
	groupVersionResource := metav1.GroupVersionResource{Group: "", Resource: "pods", Version: "v1"}

	if review.Request.Resource != groupVersionResource {
		err := admitErroNew(fmt.Errorf("expexted resource to be %s", groupVersionResource))
		klog.Error(err)
		return err.toAdmitResponse()
	}
	raw := review.Request.Object.Raw
	pod := corev1.Pod{}

	des := webhook.Codecs.UniversalDeserializer()
	if _, _, err := des.Decode(raw, nil, &pod); err != nil {
		klog.Error(err)
		return admitErroNew(err).toAdmitResponse()
	}
	reviewResponse := v1.AdmissionResponse{}
	reviewResponse.Allowed = true

	// Ignore if present ignore flag on pod labels
	if hasIgnoreNotLatestImageLabels(pod.Labels) {
		reviewResponse.Result = &metav1.Status{Message: "Ignore not latest image policy"}
	} else {
		msg, ok := haveSomeImageLatestTag(pod.Spec.Containers)
		reviewResponse.Allowed = !ok
		reviewResponse.Result = &metav1.Status{Message: msg}
	}

	return &reviewResponse
}
