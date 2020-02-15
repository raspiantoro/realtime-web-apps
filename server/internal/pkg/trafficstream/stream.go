package trafficstream

import "sync"

var (
	msgChan []chan Message
	doOnce  sync.Once
)

func init() {
	doOnce.Do(func() {
		msgChan = []chan Message{}
	})
}

//AppendChan add new message channel to msgChan
func AppendChan(ch chan Message) {
	msgChan = append(msgChan, ch)
}

//GetChan get msgChan
func GetChan() []chan Message {
	return msgChan
}
