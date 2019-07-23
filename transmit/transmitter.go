package transmit

type Transmitter interface {
	Transmit(
		data []byte,
		fileName string,
		path string,
	) error

	GetName() string
}
