package o

import (
	"fmt"
)

type Throwable interface {
	fmt.Stringer
}

type Exception struct {
	message string
}

func NewException(message string) *Exception {
	return &Exception{
		message: message,
	}
}

func (this *Exception) String() string {
	return "[Exception] " + this.message
}
