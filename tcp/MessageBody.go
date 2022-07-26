package tcp

type MessageBody struct {
	User    string
	Message string
}

func NewMessage(user string, message string) MessageBody {
	return MessageBody{User: user, Message: message}
}

func NewEmptyMessage() MessageBody {
	return MessageBody{}
}
