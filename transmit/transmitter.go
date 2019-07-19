package transmit

type Transmitter interface {
	Stor(
		data []byte,
		fileName string,
	) bool

	GetName() string
}
