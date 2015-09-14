package o_test

import (
	"fmt"
	"testing"
)

type Ctrl struct {
}

func (this *Ctrl) Actions() map[string]func() {
	return map[string]func(){
		"login": func() {
			fmt.Println("200 login")
		},
		"logout": func() {
			fmt.Println("200 logout")
		},
		"info": func() {
			fmt.Println("200 info")
		},
	}
}

func (this *Ctrl) NotFound() {
	fmt.Println("404 not found")
}

func Handler(ctrl *Ctrl, arg string) {
	if f, has := ctrl.Actions()[arg]; !has {
		ctrl.NotFound()
	} else {
		f()
	}
}

func TestMapFunc(t *testing.T) {
	m := "login"

	ctrl := new(Ctrl)
	Handler(ctrl, m)
}
