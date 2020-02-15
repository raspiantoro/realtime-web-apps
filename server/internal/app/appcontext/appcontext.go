package appcontext

import (
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/consumer"
	"github.com/raspiantoro/realtime-web-apps/server/internal/app/service"
)

//Service hold all application service instance
type Service struct {
	Traffic service.TrafficService
}

//Consumer hold all application consumer instance
type Consumer struct {
	Traffic consumer.Traffic
}

//InitService instantiate all application servic
func InitService() Service {
	svc := Service{
		Traffic: service.TrafficService{},
	}

	return svc
}

//InitConsumer instantiate all application consumer
func InitConsumer() Consumer {
	c := Consumer{
		Traffic: consumer.NewConsumer(),
	}

	return c
}
