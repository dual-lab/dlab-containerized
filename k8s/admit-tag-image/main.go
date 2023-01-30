package main

import (
	"github.com/dual-lab/dlab-containerized/k8s/admit-tag-image/pkg/net"
	"github.com/urfave/cli/v2"
	"k8s.io/klog/v2"
	"os"
)

func main() {
	var certFile string
	var certKey string
	var host string
	var port int
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "certFile",
				Value:       "",
				Usage:       "Specify cert file path",
				Destination: &certFile,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "certKey",
				Value:       "",
				Usage:       "Specify cert key file path",
				Destination: &certKey,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "host",
				Value:       "localhost",
				Usage:       "Specify service host",
				Destination: &host,
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       443,
				Usage:       "Specify server port",
				Destination: &port,
			},
		},
		Name:  "adc-notlt",
		Usage: "Admit controller tag image not latest policy",
		Action: func(context *cli.Context) error {
			net.Register()
			return net.Run(host, port, certFile, certKey)
		},
	}
	if err := app.Run(os.Args); err != nil {
		klog.Fatal(err)
	}
}
