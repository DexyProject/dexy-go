package subscribers

type Subscriber interface {
	Subscribe() error
	Listen() (string, error)
}