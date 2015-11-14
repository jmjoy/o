package o

type StructBase struct {
	name string
	age  int
}

func (this *StructBase) Name(namex ...string) (name string) {
	// get
	if len(namex) == 0 {
		return this.name
	}

	// set
	this.name = namex[0]
	return
}

func (this *StructBase) Age(agex ...int) (age int) {
	// get
	if len(agex) == 0 {
		return this.age
	}

	// set
	this.age = agex[0]
	return
}

// ~~~~(>_<)~~~~
func (this *StructBase) Field(field string, vx ...interface{}) (value interface{}) {
	//v := reflect.ValueOf(this).Elem().FieldByName(field)

	// get
	// reflect.Value.Interface: cannot return value obtained from unexported field or method
	//if len(vx) == 0 {
	//return v.Interface()
	//}

	// set
	//reflect: reflect.Value.Set using value obtained using unexported field
	//v.Set(reflect.ValueOf(vx[0]))
	return
}
