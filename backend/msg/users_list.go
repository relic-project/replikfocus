package msg

type ConnectedUsersMsg struct {
	Message
	Users []string `json:"users"`
}

func NewConnectedUsersMsg(users []string) ConnectedUsersMsg {
	return ConnectedUsersMsg{
		Message: Message{
			Type: "connected_users",
		},
		Users: users,
	}
}
