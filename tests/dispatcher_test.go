package tests

import (
	"context"
	"github.com/lapuda/data_dispatcher"
	"log"
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	dispatcher := data_dispatcher.NewDispatcherAndRun(func(data []byte) {
		log.Printf("recive data %v\n", string(data))
	}, ctx)
	dispatcher.Collect([]byte("test"))
}
