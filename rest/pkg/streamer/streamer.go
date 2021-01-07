package streamer

import (
	"github.com/raspiantoro/realtime-web-apps/rest/pkg/streamer/natsstreaming"
)

//Streamer define interface for streamer
type Streamer interface {
	Emit(topic string, payload []byte) (err error)
	Listen(topic string, durable string, handler natsstreaming.Callback) (listenerValue natsstreaming.ListenerValue, err error)
	Close() (err error)
}

//Option define struct option for streamer
type Option struct {
	Streamer   string
	StanOption natsstreaming.ClientOption
}

//NewStreamer create new streamer instance
func NewStreamer(opt Option) (streamer Streamer, err error) {
	if opt.Streamer == StanStreamerType {
		streamer, err = natsstreaming.NewNatsStreamingClient(opt.StanOption)
		return
	}

	return
}
