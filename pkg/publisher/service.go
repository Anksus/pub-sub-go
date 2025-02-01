package publisher

import "anksus/pub-sub-go/pkg/topic"

type Service interface {
	RegisterTopic(t *topic.Topic)
	DeregisterTopic(t *topic.Topic)
	Publish(t *topic.Topic, m string) error
}
