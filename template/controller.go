package main

import (
	"time"
)

type controller struct {
	interval         time.Duration
	kube *Kube
}

func newController(kube *Kube, interval time.Duration) *controller {
	c := &controller{
		interval:         interval,
		kube: kube,
	}
	return c
}

func (c *controller) Run(stopCh chan struct{}) {
	log().Info("Starting controller")

	<-stopCh
	log().Info("Stopping controller")
}
