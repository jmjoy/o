package o_test

import (
	"testing"

	"github.com/jmjoy/o"
)

func TestStructBase(t *testing.T) {
	base := new(o.StructBase)

	base.Name("helloworld")
	if base.Name() != "helloworld" {
		t.Fail()
	}

	base.Age(22)
	if base.Age() != 22 {
		t.Fail()
	}

	//base.Field("name", "gopher")
	//if base.Field("name").(string) != "gopher" {
	//    t.Fail()
	//}

	//base.Field("age", 100)
	//if base.Field("age").(int) != 100 {
	//    t.Fail()
	//}
}
