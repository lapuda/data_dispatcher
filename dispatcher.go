package data_dispatcher

import (
	"context"
	"github.com/lapuda/data_dispatcher/core"
	"log"
)
import event "github.com/lapuda/event_center/core"

type Dispatcher struct {
	Ctx         context.Context
	eventCenter *event.EventCenter
}

func (d *Dispatcher) initEnv() {
	d.eventCenter = event.CreateEventCenter(d.Ctx)
	d.eventCenter.Register(core.CollectEvent{}.Name())
}

func (d *Dispatcher) AddTarget(targetName string, writeFunc event.EventHandler) {
	d.eventCenter.Subscribe(core.CollectEvent{}.Name(), event.HandlerName(targetName), writeFunc)
}

func (d *Dispatcher) Collect(data interface{}) {
	error := d.eventCenter.SendEvent(core.CollectEvent{CollectData: data})
	if error != nil {
		log.Panicf("during collect data occer error: %v \n", error)
	}
}

func (d *Dispatcher) Run() {
	d.initEnv()
}

func NewDispatcherAndRun(ctx context.Context) *Dispatcher {
	dispatcher := Dispatcher{Ctx: ctx}
	dispatcher.Run()
	return &dispatcher
}
