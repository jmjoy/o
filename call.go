package o

import (
	"fmt"
)

// 增强抽象能力

type CallStruct struct {
	fieldA string
	fieldB int
	fieldC []float32
	fieldD map[string]interface{}
}

func (this *CallStruct) Get(field string) (value interface{}, err error) {
	switch field {
	case "fieldA":
		value = this.fieldA

	case "fieldB":
		value = this.fieldB

	case "fieldC":
		value = this.fieldC

	case "fieldD":
		value = this.fieldD

	default:
		err = fmt.Errorf("Field not found: %s in %s", field, "CallStruct")
	}
	return
}

func (this *CallStruct) MustGet(field string) (value interface{}) {
	value, err := this.Get(field)
	if err != nil {
		panic(err)
	}
	return
}

func (this *CallStruct) Set(field string, value interface{}) (err error) {
	switch field {
	case "fieldA":
		if v, ok := value.(string); ok {
			this.fieldA = v
			return
		}
	case "fieldB":
		if v, ok := value.(int); ok {
			this.fieldB = v
			return
		}
	case "fieldC":
		if v, ok := value.([]float32); ok {
			this.fieldC = v
			return
		}
	case "fieldD":
		if v, ok := value.(map[string]interface{}); ok {
			this.fieldD = v
			return
		}
	}
	return fmt.Errorf("Field not found: %s in %s", field, "CallStruct")
}

func (this *CallStruct) MustSet(field string, value interface{}) {
	if err := this.Set(field, value); err != nil {
		panic(err)
	}
}
