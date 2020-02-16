package appcontext

import (
	"context"

	"github.com/raspiantoro/realtime-web-apps/server/internal/app/consumer"
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/service"
	streamPB "github.com/raspiantoro/realtime-web-apps/server/pkg/stream/v1"
)

//Service hold all application service instance
type Service struct {
	Stream streamPB.StreamServer
}

//Consumer hold all application consumer instance
type Consumer struct {
	Stream consumer.StreamConsumer
}

//InitService instantiate all application service
func InitService(ctx context.Context) Service {
	svc := Service{
		Stream: service.NewStreamService(ctx),
	}

	return svc
}

//InitConsumer instantiate all application consumer
func InitConsumer() Consumer {
	c := Consumer{
		Stream: consumer.NewStreamConsumer(),
	}

	return c
}
