package o

type single struct {
}

func (*single) Do() {
}

var singleInstance *single

func InstanceSingle() *single {
	if singleInstance == nil {
		singleInstance = new(single)
	}
	return singleInstance
}
