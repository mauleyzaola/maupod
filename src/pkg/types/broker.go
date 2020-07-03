package types

type Broker interface {
	Register() error
	Close()
}
