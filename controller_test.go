// 面向组合的设计
package o_test

import (
	"io"
	"testing"
)

// Context
type Context struct {
	in  io.Reader
	out io.Writer
}

// Actions Component
type ActionsComponent struct {
	actions map[string]func()
}

func (this *ActionsComponent) SetActions(actions map[string]func()) {
	this.actions = actions
}

func (this *ActionsComponent) Actions() map[string]func() {
	return this.actions
}

// RESTful Component
type RESTfulComponent struct {
}

func (this *RESTfulComponent) doGet() {
}

func (this *RESTfulComponent) doPost() {
}

func (this *RESTfulComponent) doPut() {
}

func (this *RESTfulComponent) doDelete() {
}

func (this *RESTfulComponent) doContinue() {
}

func (this *RESTfulComponent) doHead() {
}

func (this *RESTfulComponent) doOptions() {
}

func (this *RESTfulComponent) doPatch() {
}

type ActionsController struct {
	*Context
	*ActionsComponent
}

func NewActionsController() *ActionsController {
	this := new(ActionsController)

	this.SetActions(map[string]func(){
		"login": func() {

		},
	})

	return this
}

type RESTfulController struct {
	*Context
	*RESTfulComponent
}

func TestController(t *testing.T) {

}
