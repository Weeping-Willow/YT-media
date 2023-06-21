package main

import (
	"os"
	"os/signal"
	"syscall"

	"media-server/app"

	"github.com/sirupsen/logrus"
)

func main() {
	application := app.NewApp()

	if err := application.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start application")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	application.Stop()
}
