package providers

type Provider interface {
	Stor(
		data []byte,
		fileName string,
	) bool

	GetName() string
}
