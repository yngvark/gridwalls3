package network

type Listener interface {
	NotifyListeneres(msg string)
}

type Broadcaster struct {
	listeners []Listener
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		listeners: make([]Listener, 0),
	}
}

func (n *Broadcaster) AddListener(l Listener) {
	n.listeners = append(n.listeners, l)
}

func (n *Broadcaster) NotifyListeneres(msg string) {
	for _, l := range n.listeners {
		l.NotifyListeneres(msg)
	}
}

type MessageSender interface {
	SendMsg(msg string) error
}
