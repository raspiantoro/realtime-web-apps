package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/stream"
	"github.com/raspiantoro/realtime-web-apps/worker/pkg/streamer"
	"github.com/raspiantoro/realtime-web-apps/worker/pkg/streamer/natsstreaming"
)

//StreamConsumer define struct for stream consumer
type StreamConsumer struct {
	streamer streamer.Streamer
}

//NewStreamConsumer create stream consumer instance
func NewStreamConsumer(streamer streamer.Streamer) StreamConsumer {
	return StreamConsumer{
		streamer: streamer,
	}
}

//Start starting consumer
func (c *StreamConsumer) Start(ctx context.Context) {

	sub, err := c.streamer.Listen("streamdata", "", sendMessage)
	if err != nil {
		log.Printf("Error streamer listener: %s", err)
	}

	go func() {
		select {
		case <-ctx.Done():
			log.Print("unsubscribe streamer topic")
			sub.Unsubscribe()
		}
	}()

}

func sendMessage(data []byte, ack natsstreaming.Ack) (err error) {
	ack()
	msg := stream.Message{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return
	}

	log.Printf("receive new message from streamer")

	streamChan := stream.GetChan()

	for key, ch := range streamChan {
		log.Printf("send message to streamchan: %d", key)
		ch <- msg
	}

	return
}
