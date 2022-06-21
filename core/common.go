package core

import event "github.com/lapuda/event_center/core"

type CollectEvent struct {
	CollectData []byte
}

func (c CollectEvent) Name() event.EventName {
	return event.EventName("CollectEvent")
}

func (c CollectEvent) Data() interface{} {
	return c.CollectData
}

type WriteFunc func(data []byte)
