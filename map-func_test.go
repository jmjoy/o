package o_test

import (
	"fmt"
	"reflect"
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

func _TestMapFunc(t *testing.T) {
	m := "login"

	ctrl := new(Ctrl)
	Handler(ctrl, m)
}

func MapStructMethods(s interface{}) (map[string]func(), error) {
	value := reflect.ValueOf(s)
	numMethod := value.NumMethod()
	if numMethod == 0 {
		panic("No methods!")
	}
	for i := 0; i < numMethod; i++ {
		fmt.Println(value.Method(i))
	}
	return nil, nil
}

type AStruct struct {
	AField string
}

func (this *AStruct) AFunc() {
	fmt.Println("Hello A")
}

func (this *AStruct) BFunc(i int, s string) {
	fmt.Printf("Hello %d %s\n", i, s)
}

func TestMapStructFuncs(t *testing.T) {
	//MapStructMethods(&AStruct{})
}

type Controller struct {
	actions map[string]func()
	Name    string
}

func NewController() *Controller {
	this := new(Controller)
	this.Name = "begin"
	this.InitFuncs()
	return this
}

func (this *Controller) InitFuncs() {
	this.actions = make(map[string]func())

	this.actions["login"] = func() {
		fmt.Println(this.Name)
	}

	this.actions["login"] = func() {
		fmt.Println(this.Name)
	}
}
