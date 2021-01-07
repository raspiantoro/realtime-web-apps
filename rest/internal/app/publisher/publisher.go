package publisher

import (
	"encoding/json"
	"log"

	"github.com/raspiantoro/realtime-web-apps/rest/internal/app/message"
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer"
)

//Publisher define Publisher interface
type Publisher interface {
	Publish(message.StreamData) (err error)
}

//publisher define publisher struct
type publisher struct {
	stmr streamer.Streamer
}

//NewPublisher create new publisher instance
func NewPublisher(stmr streamer.Streamer) Publisher {
	return &publisher{
		stmr: stmr,
	}
}

func (p *publisher) Publish(msg message.StreamData) (err error) {
	byteMessage, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error on marshaling strMsg: %s", err)
	}

	err = p.stmr.Emit("streamdata", byteMessage)
	return
}
