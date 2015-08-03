package o_test

import (
	"testing"

	"github.com/jmjoy/o"
)

func TestSingle(t *testing.T) {
	single := o.InstanceSingle()
	single.Do()
}
