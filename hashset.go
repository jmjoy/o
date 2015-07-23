package o

type HashSet struct {
	m    map[interface{}]struct{}
	size int
}

func NewHashSet() *HashSet {
	return NewHashSetCap(10)
}

func NewHashSetCap(size int) *HashSet {
	hashset := new(HashSet)
	hashset.m = make(map[interface{}]struct{}, size)
	return hashset
}

func (this *HashSet) Add(i interface{}) {
	this.m[i] = struct{}{}
	this.size++
}

func (this *HashSet) Remove(i interface{}) {
	delete(this.m, i)
	this.size--
}

func (this *HashSet) Has(i interface{}) bool {
	_, ok := this.m[i]
	return ok
}

func (this *HashSet) Size() int {
	return this.size
}

func (this *HashSet) Walk(f func(interface{}) error) error {
	for i := range this.m {
		err := f(i)
		if err != nil {
			return err
		}
	}
	return nil
}
