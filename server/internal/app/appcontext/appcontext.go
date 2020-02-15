package appcontext

import "github.com/raspiantoro/realtime-web-apps/server/internal/app/service"

//Service hold all service for appcontext
type Service struct {
	Traffic service.TrafficService
}

//InitService instantiate service fot appcontext
func InitService() Service {
	svc := Service{
		Traffic: service.TrafficService{},
	}

	return svc
}
