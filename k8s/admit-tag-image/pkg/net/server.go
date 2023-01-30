package net

import (
	"fmt"
	"github.com/dual-lab/admit-webook-boilerplate/pkg/webhook/net"
	"net/http"
)

func Run(host string, port int, certFile, keyFile string) error {
	tslconfig := net.TLSConfig{
		CertFile: certFile, KeyFile: keyFile,
	}
	server := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", host, port),
		TLSConfig: tslconfig.Load(),
	}
	return server.ListenAndServeTLS("", "")
}

func Register() {
	http.Handle("/pods", podAdmitFuncGet())
	http.Handle("/podtemplates", podTemplatesFuncGet())
	http.Handle("/readyz", http.HandlerFunc(net.Readyz))
}
