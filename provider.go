package mobStor

type provider interface {
	stor(
		data []byte,
		fileName string,
	) bool

	getName() string
}
