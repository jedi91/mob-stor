package transmit

type Transmitter interface {
	Transmit(
		data []byte,
		filePath string,
		containerName string,
	) error

	GetName() string
}
