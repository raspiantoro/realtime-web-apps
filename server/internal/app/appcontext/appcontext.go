package appcontext

import (
	"context"
	"log"

	"github.com/raspiantoro/realtime-web-apps/server/internal/app/consumer"
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/service"
	streamPB "github.com/raspiantoro/realtime-web-apps/server/pkg/stream/v1"
	"github.com/raspiantoro/realtime-web-apps/worker/pkg/streamer"
)

//Service hold all application service instance
type Service struct {
	Stream streamPB.StreamServer
}

//Consumer hold all application consumer instance
type Consumer struct {
	Stream consumer.StreamConsumer
}

//Streamer hold all streamer application context
type Streamer struct {
	Stan streamer.Streamer
}

//StreamerOption hold all streamer options application context
type StreamerOption struct {
	Stan streamer.Option
}

//InitService instantiate all application service
func InitService(ctx context.Context) Service {
	svc := Service{
		Stream: service.NewStreamService(ctx),
	}

	return svc
}

//InitConsumer instantiate all application consumer
func InitConsumer(streamer streamer.Streamer) Consumer {
	c := Consumer{
		Stream: consumer.NewStreamConsumer(streamer),
	}

	return c
}

//InitStreamer instantiate all streamer instance
func InitStreamer(opt StreamerOption) (stmr Streamer) {
	stan, err := streamer.NewStreamer(opt.Stan)
	if err != nil {
		log.Fatalf("failed to connect to streamer server: %s", err)
	}

	stmr = Streamer{
		Stan: stan,
	}
	return
}
