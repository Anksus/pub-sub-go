package subscriber

type Service interface {
	Notify(m string)
}
