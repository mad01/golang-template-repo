package main

import (
	"github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}/log"
)

type Event struct {
	Name string
}

func newEvent(serviceName, manifest string) *Event {
	e := &Event{
		Name:     serviceName,
	}
	return e
}

type EventHandler struct {
	events chan Event
}

func newEventHandler() *EventHandler {
	e := &EventHandler{
		events: make(chan Event, 100),
	}
	return e
}

func (e *EventHandler) AddEvent(event *Event) {
	e.events <- *event
}

func (e *EventHandler) Run(stopCh chan struct{}) {
	e.worker(stopCh)
}

func (e *EventHandler) worker(stopCh chan struct{}) {
	for {
		select {
		case event := <-e.events:
			log.Infof("%v", event)
		case _ = <-stopCh:
			return

		}
	}
}
