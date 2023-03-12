package msg

type PlaySoundMsg struct {
	Message
	Sound string `json:"sound"`
}

func NewPlaySoundMsg(sound string) PlaySoundMsg {
	return PlaySoundMsg{
		Message: Message{
			Type: "play_sound",
		},
		Sound: sound,
	}
}
