package o2_test

import (
	"fmt"
	"testing"
)

type IActions interface {
	Actions() map[string]func()
}

type Controller struct {
	Name string
}

func NewController(name string) *Controller {
	return &Controller{
		Name: name,
	}
}

func (this *Controller) Actions() map[string]func() {
	fmt.Println("Access Actions")

	return map[string]func(){
		"login": func() {
			fmt.Println(this.Name + ": login")
		},

		"logout": func() {
			fmt.Println(this.Name + ": Logout")
		},
	}
}

var globalFuncMap = make(map[IActions]map[string]func())

func HandleActions(actions IActions, action string) {
	//typeKey := reflect.TypeOf(actions).String()
	//if typeKey == "" {
	//    panic("Empty type")
	//}

	funcMap, has := globalFuncMap[actions]
	if !has {
		funcMap = actions.Actions()
		globalFuncMap[actions] = funcMap
	}

	fn, has := funcMap[action]
	if !has {
		panic("No this function")
	}

	fn()

	//delete(globalFuncMap, actions)
}

func TestMe(t *testing.T) {
	a := NewController("A")
	HandleActions(a, "login")
	HandleActions(a, "logout")
	HandleActions(NewController("B"), "logout")
}
