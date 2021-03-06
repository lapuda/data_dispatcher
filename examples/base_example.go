package main

import (
	"context"
	"github.com/lapuda/data_dispatcher"
	"log"
	"sync"
)

var wg = sync.WaitGroup{}

func PrintDispatcherWrite(data interface{}) {
	log.Printf("学校的输出 : %v \n", data)
	wg.Done()
}

func PrintDispatcherWriteAndHello(data interface{}) {
	log.Printf("银行的输出 : %v \n", data)
	wg.Done()
}

func main() {
	wg.Add(6)
	ctx := context.Context(context.Background())
	dispatcher := data_dispatcher.NewDispatcherAndRun(ctx)

	dispatcher.AddTarget("学校", PrintDispatcherWrite)
	dispatcher.AddTarget("银行", PrintDispatcherWriteAndHello)

	dispatcher.Collect("hello")
	dispatcher.Collect("hello2")
	dispatcher.Collect("hello3")
	// only test
	wg.Wait()
}
