package signals

import (
	"os"
	"syscall"
)

var shutdownSIGN = []os.Signal{os.Interrupt, syscall.SIGTERM}
