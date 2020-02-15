package consumer

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/trafficstream"
)

//Traffic define struct for traffic consumer
type Traffic struct{}

//NewConsumer create consumer instance
func NewConsumer() Traffic {
	return Traffic{}
}

//Start starting consumer
func (c *Traffic) Start(ctx context.Context) {
	var streamChan map[int64]chan trafficstream.Message

	go func() {
		for {
			time.Sleep(1 * time.Second)

			min := 300
			max := 1000

			msg := trafficstream.Message{
				TrafficCount: rand.Intn(max-min) + min,
			}

			streamChan = trafficstream.GetChan()

			for key, ch := range streamChan {
				log.Printf("send message to streamchan to chan: %d", key)
				ch <- msg
			}

		}
	}()

}
