package subscriber

import "fmt"

type Subscriber struct {
	Nane string
}

func (s *Subscriber) Notify(m string) {
	fmt.Println(fmt.Sprintf("Message received on subscriber %s: %s", s.Nane, m))

}
