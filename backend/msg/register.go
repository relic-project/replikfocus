package msg

type RegisterMsg struct {
	Message
	Username string `json:"username"`
}

func NewRegisterMsg(username string) RegisterMsg {
	return RegisterMsg{
		Message: Message{
			Type: "register",
		},
		Username: username,
	}
}
