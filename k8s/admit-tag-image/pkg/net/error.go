package net

import (
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type admitError string

func admitErroNew(err error) admitError  {
  return admitError(err.Error())
}

func (receiver admitError) Error() string {
	return string(receiver)
}

func (receiver admitError) toAdmitResponse() *v1.AdmissionResponse {
	return &v1.AdmissionResponse{
		Result: &metav1.Status{
			Message: receiver.Error(),
		},
	}
}
