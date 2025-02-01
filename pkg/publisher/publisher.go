package publisher

import (
	"anksus/pub-sub-go/pkg/topic"
)

type Publisher struct {
	Topics []*topic.Topic
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) RegisterTopic(t *topic.Topic) {
	p.Topics = append(p.Topics, t)
}

func (p *Publisher) DeregisterTopic(t *topic.Topic) {
	p.Topics = removeEleInPlace(p.Topics, t)
}

func (p *Publisher) Publish(m string) {
	for _, t := range p.Topics {
		for _, s := range t.Subscribers {
			s.Notify(m)
		}
	}
}

func removeEleInPlace(t []*topic.Topic, tr *topic.Topic) []*topic.Topic {
	j := 0
	for _, v := range t {
		if v != tr {
			t[j] = v
			j++
		}
	}
	return t[:j]
}

func removeElement(t []*topic.Topic, tr *topic.Topic) []*topic.Topic {
	for i, v := range t {
		if tr == v {
			return append(t[:i], t[:i+1]...)
		}
	}
	return t
}
