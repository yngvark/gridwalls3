package pubsub

type Publisher interface {
	SendMsg(msg string) error
}
