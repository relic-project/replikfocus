package msg

import "time"

type NewTimerMsg struct {
	Message
	Mode   string    `json:"mode"`
	Expire time.Time `json:"expire"`
}

func NewNewTimerMsg(mode string, expire time.Time) NewTimerMsg {
	return NewTimerMsg{
		Message: Message{
			Type: "new_timer",
		},
		Mode:   mode,
		Expire: expire,
	}
}
