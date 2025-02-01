package main

import (
	"anksus/pub-sub-go/pkg/publisher"
	"anksus/pub-sub-go/pkg/subscriber"
	"anksus/pub-sub-go/pkg/topic"
)

func main() {

	s1 := &subscriber.Subscriber{
		Nane: "s1",
	}

	s2 := &subscriber.Subscriber{
		Nane: "s2",
	}

	// topics
	t1 := &topic.Topic{}
	t1.RegisterSubscriber(s1)
	t1.RegisterSubscriber(s2)

	t2 := &topic.Topic{}

	p1 := publisher.NewPublisher()
	p1.RegisterTopic(t1)

	p2 := publisher.NewPublisher()
	p2.RegisterTopic(t2)
	p2.RegisterTopic(t1)

	p1.Publish("hello")

}
