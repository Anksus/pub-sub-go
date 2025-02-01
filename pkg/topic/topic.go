package topic

import "anksus/pub-sub-go/pkg/subscriber"

type Topic struct {
	Subscribers []subscriber.Service
}

func (t *Topic) RegisterSubscriber(s subscriber.Service) {
	t.Subscribers = append(t.Subscribers, s)
}
