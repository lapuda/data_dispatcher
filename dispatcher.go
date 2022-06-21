package data_dispatcher

import (
	"context"
	"github.com/lapuda/data_dispatcher/core"
	"log"
)
import event "github.com/lapuda/event_center/core"

const SubscribeDispatchHandler = "Subscribe_Dispatch_Handler"

type Dispatcher struct {
	WriteFunc   core.WriteFunc
	Ctx         context.Context
	eventCenter *event.EventCenter
}

func (d *Dispatcher) initEnv() {
	collectEvent := core.CollectEvent{}
	d.eventCenter = event.CreateEventCenter(d.Ctx)

	//register
	d.eventCenter.Register(collectEvent.Name())
	// disptch process
	d.eventCenter.Subscribe(collectEvent.Name(), SubscribeDispatchHandler, d.dispatch)
}

func (d *Dispatcher) dispatch(param interface{}) {
	d.WriteFunc(param.([]byte))
}

func (d *Dispatcher) Collect(data []byte) {
	error := d.eventCenter.SendEvent(core.CollectEvent{CollectData: data})
	if error != nil {
		log.Panicf("during collect data occer error: %v \n", error)
	}
}

func (d *Dispatcher) Run() {
	d.initEnv()
}

func NewDispatcherAndRun(writeFunc core.WriteFunc, ctx context.Context) *Dispatcher {
	dispatcher := Dispatcher{WriteFunc: writeFunc, Ctx: ctx}
	dispatcher.Run()
	return &dispatcher
}
