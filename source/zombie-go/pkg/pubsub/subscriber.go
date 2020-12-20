package pubsub

type Subscriber interface {
	MsgReceived(msg string)
}
