package signals

import "os"

var shutdownSIGN = []os.Signal{os.Interrupt}
