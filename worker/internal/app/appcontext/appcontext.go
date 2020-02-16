package appcontext

import (
	"log"

	"github.com/raspiantoro/realtime-web-apps/worker/internal/app/publisher"
	"github.com/raspiantoro/realtime-web-apps/worker/internal/app/task"
	"github.com/raspiantoro/realtime-web-apps/worker/internal/pkg/streamer"
)

//Streamer hold all streamer application context
type Streamer struct {
	Stan streamer.Streamer
}

//Publisher hold all publisher application context
type Publisher struct {
	StreamPublisher publisher.Publisher
}

//Task hold all task application context
type Task struct {
	StreamTask task.Task
}

//InitStanStreamer instantiate all streamer instance
func InitStanStreamer(opt streamer.Option) (stan streamer.Streamer) {
	stan, err := streamer.NewStreamer(opt)
	if err != nil {
		log.Fatalf("failed to connect to streamer server: %s", err)
	}

	return
}

//InitPublisher instantiate all publisher instance
func InitPublisher(stmr streamer.Streamer) (pubs Publisher) {
	return Publisher{
		StreamPublisher: publisher.NewPublisher(stmr),
	}

}

//InitTask instantiate all task instance
func InitTask(pubs Publisher) (tsk Task) {
	return Task{
		StreamTask: task.NewStreamTask(pubs.StreamPublisher),
	}
}
