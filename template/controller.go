package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type controller struct {
	interval   time.Duration
	kube       *Kube
	stopCh     chan struct{}
	signalChan chan os.Signal
}

func newController(kube *Kube, interval time.Duration) *controller {
	c := &controller{
		interval:   interval,
		kube:       kube,
		stopCh:     make(chan struct{}),
		signalChan: make(chan os.Signal, 1),
	}
	return c
}

func (c *controller) Run() {
	log().Info("Starting controller")

	go c.worker(c.stopCh)

	signal.Notify(c.signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-c.signalChan // Block until sigterm
	defer close(c.stopCh)

	log().Info("Stopping controller")
	return
}

func (c *controller) worker(stopCh chan struct{}) {
	for {
		select {
		default:
			// TODO: do a bit of the work
		case <-stopCh:
			log().Info("Stopping worker since stopCh closed")
			return
		}
	}
}
