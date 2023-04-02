package main

import (
	"github.com/dual-lab/dlab-containerized/dns/duck-dns/operator/pkg/signals"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	//TODO: parse flag from command line

	ctx := signals.SetupSignalHandler()
	logger := klog.FromContext(ctx)

	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Error(err, "Error on get configuration from kube cluster")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err, "Error building kube client")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	logger.Info("TODO", kubeClient)

	//TODO create standard informer
	//TODO create custom controller clinet
	//TODO create custom informer
	//TODO create new controller
	//TODO start informers
	//TODO start controller and loop
}
