package o

type single struct {
}

func (*single) Do() {
}

var singleInstance single

func InstanceSingle() *single {

}
