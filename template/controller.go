package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type controller struct {
	interval time.Duration
	kube     *Kube
}

func newController(kube *Kube, interval time.Duration) *controller {
	c := &controller{
		interval: interval,
		kube:     kube,
	}
	return c
}

func (c *controller) Run() {
	log().Info("Starting controller")

	go c.worker()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log().Info("Stopping controller")
}

func (c *controller) worker() {
	for {
	}
}
