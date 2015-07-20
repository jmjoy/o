package o

type Throwable interface {
	Message() string
}

type Exception struct {
	message string
}

func NewException(message string) *Exception {
	return &Exception{
		message: message,
	}
}

func (this *Exception) Message() string {
	return this.message
}
