package main

import (
	"context"
	"github.com/lapuda/data_dispatcher"
	"log"
	"time"
)

func PrintDispatcherWrite(data []byte) {
	log.Printf("print output : %v \n", string(data))
}

func main() {
	ctx := context.Context(context.Background())
	dispatcher := data_dispatcher.NewDispatcherAndRun(PrintDispatcherWrite, ctx)

	dispatcher.Collect([]byte("hello"))
	dispatcher.Collect([]byte("hello2"))
	dispatcher.Collect([]byte("hello3"))
	// only test
	time.Sleep(1 * time.Second)
}
