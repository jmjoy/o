package o

type HashSet map[interface{}]struct{}

func NewHashSet() HashSet {
	return NewHashSetCap(1)
}

func NewHashSetCap(size int) HashSet {
	return make(map[interface{}]struct{}, size)
}

func (this HashSet) Add(i interface{}) {
	this[i] = struct{}{}
}

func (this HashSet) Remove(i interface{}) {
	delete(this, i)
}

func (this HashSet) Has(i interface{}) bool {
	_, ok := this[i]
	return ok
}

func (this HashSet) Size() int {
	return len(this)
}

func (this HashSet) Walk(f func(interface{}) error) error {
	for i := range this {
		err := f(i)
		if err != nil {
			return err
		}
	}
	return nil
}
