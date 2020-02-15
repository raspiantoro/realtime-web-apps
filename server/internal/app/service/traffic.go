package service

import (
	"log"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/trafficstream"
	trafficPB "github.com/raspiantoro/realtime-web-apps/server/pkg/traffic/v1"
)

type TrafficService struct{}

//NewTrafficService create new trafficService that implement trafficPB.TrafficServer
func NewTrafficService() trafficPB.TrafficServer {
	return &TrafficService{}
}

//GetTrafficCount service for handling GetTrafficCount request
func (s *TrafficService) GetTrafficCount(trafficCount *trafficPB.TrafficCountRequest, stream trafficPB.Traffic_GetTrafficCountServer) error {
	ch := make(chan trafficstream.Message, 1)
	trafficstream.AppendChan(ch)

	select {
	case <-stream.Context().Done():
		log.Printf("Process has been canceled")
		close(ch)
	case val := <-ch:
		resp := &trafficPB.TrafficCountResponse{
			TrafficCount: val.TrafficCount,
		}

		stream.Send(resp)
	}

	return nil
}
