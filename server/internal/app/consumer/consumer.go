package consumer

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/stream"
)

//StreamConsumer define struct for stream consumer
type StreamConsumer struct{}

//NewStreamConsumer create stream consumer instance
func NewStreamConsumer() StreamConsumer {
	return StreamConsumer{}
}

//Start starting consumer
func (c *StreamConsumer) Start(ctx context.Context) {
	var streamChan map[int64]chan stream.Message

	go func() {
		for {
			time.Sleep(1 * time.Second)

			min := 300
			max := 1000

			msg := stream.Message{
				Count: rand.Intn(max-min) + min,
			}

			streamChan = stream.GetChan()

			for key, ch := range streamChan {
				log.Printf("send message to streamchan: %d", key)
				ch <- msg
			}

		}
	}()

}
