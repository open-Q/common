package golang

import (
	"os"
	"os/signal"
	"syscall"
)

// InterruptHook starts new interrupt hook.
func InterruptHook(hFunc func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	if hFunc != nil {
		hFunc()
	}
}
