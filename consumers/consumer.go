package consumers

type Consumer interface {
	StartConsuming() error
	StopConsuming()
}