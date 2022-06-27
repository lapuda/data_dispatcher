package core

import event "github.com/lapuda/event_center/core"

type CollectEvent struct {
	CollectData interface{}
}

func (c CollectEvent) Name() event.EventName {
	return event.EventName("CollectEvent")
}

func (c CollectEvent) Data() interface{} {
	return c.CollectData
}
