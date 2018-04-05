package main

import (
	"os"
	"os/signal"
)

// ServerInterrupt is a public channel for handling
// os interrupts so we can gracefully shut down the server
var ServerInterrupt chan os.Signal

// InitInterruptHandler initializes ServerInterrupt
func InitInterruptHandler() {
	ServerInterrupt = make(chan os.Signal, 1)
	signal.Notify(ServerInterrupt, os.Interrupt)
}
