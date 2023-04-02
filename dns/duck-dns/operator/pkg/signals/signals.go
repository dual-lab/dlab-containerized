package signals

import (
	"context"
	"os"
	"os/signal"
)

var oneSignalHandlerGuard = make(chan struct{})

func SetupSignalHandler() context.Context {
	close(oneSignalHandlerGuard)

	c := make(chan os.Signal, 2)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(c, shutdownSIGN...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1)
	}()

	return ctx
}
