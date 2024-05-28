// microservice1/handler/event_handler.go
package handler

import (
    "log"
    "microservice1/model"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
    return &EventHandler{}
}

func (h *EventHandler) HandleEvent(event *model.Event) {
    log.Printf("Received event: %+v", event)
	
    // Perform actions based on the received event
}