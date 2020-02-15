package trafficstream

import "sync"

var (
	msgChan map[int64]chan Message
	doOnce  sync.Once
)

func init() {
	doOnce.Do(func() {
		msgChan = make(map[int64]chan Message)
	})
}

//AppendChan add new message channel to msgChan
func AppendChan(key int64, ch chan Message) {
	msgChan[key] = ch
}

//GetChan get msgChan
func GetChan() map[int64]chan Message {
	return msgChan
}
