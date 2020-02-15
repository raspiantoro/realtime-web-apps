package service

import (
	"log"
	"time"

	"github.com/raspiantoro/realtime-web-apps/server/internal/pkg/trafficstream"
	trafficPB "github.com/raspiantoro/realtime-web-apps/server/pkg/traffic/v1"
)

//TrafficService define struct for traffic service
type TrafficService struct{}

//NewTrafficService create new trafficService that implement trafficPB.TrafficServer
func NewTrafficService() trafficPB.TrafficServer {
	return &TrafficService{}
}

//GetTrafficCount service for handling GetTrafficCount request
func (s *TrafficService) GetTrafficCount(trafficCount *trafficPB.TrafficCountRequest, stream trafficPB.Traffic_GetTrafficCountServer) error {
	ch := make(chan trafficstream.Message, 1)
	key := time.Now().UnixNano()
	trafficstream.AppendChan(key, ch)

	for {
		select {
		case <-stream.Context().Done():
			log.Printf("Process has been canceled")
			delete(trafficstream.GetChan(), key)
			close(ch)
			return nil
		case val := <-ch:
			resp := &trafficPB.TrafficCountResponse{
				TrafficCount: uint64(val.TrafficCount),
			}

			log.Print("receive message from consumer")
			stream.Send(resp)
		}
	}
}
