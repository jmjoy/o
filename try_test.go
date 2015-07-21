package o_test

import (
	"testing"

	"github.com/jmjoy/o"
)

func TestTry(t *testing.T) {
	exception := o.NewException("SHIT")

	o.Try(func() {
		o.Throw(exception)

	}, func(throw o.Throwable) {
		if throw.(*o.Exception) != exception {
			t.Fatal("Exception not equal!")
		}

	})
}

func TestTry2(t *testing.T) {
	exception := o.NewException("SHIT")

	o.Try(func() {

		o.Try(func() {
			o.Throw(exception)

		}, func(throw o.Throwable) {
			o.Throw(throw)

		})

	}, func(throw o.Throwable) {
		if throw.(*o.Exception) != exception {
			t.Fatal("Exception not equal!")
		}

	})
}
