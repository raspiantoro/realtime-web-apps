package task

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/message"
	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/publisher"
)

//Task define task interface
type Task interface {
	Run(ctx context.Context)
}

//StreamTask define stream task struct
type StreamTask struct {
	pub publisher.Publisher
}

//NewStreamTask create new StreamTask instance
func NewStreamTask(pub publisher.Publisher) Task {
	return &StreamTask{
		pub: pub,
	}
}

//Run executing task
func (t *StreamTask) Run(ctx context.Context) {
	waitChan := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			waitChan <- true
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Print("cancelling all running tasks")
			return
		case <-waitChan:
			min := 300
			max := 1000

			msg := message.StreamData{
				Count: rand.Intn(max-min) + min,
			}

			t.pub.Publish(msg)
			log.Print("new message has been published")
		}
	}
}
