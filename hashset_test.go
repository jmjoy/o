package o_test

import (
	"sort"
	"testing"

	"github.com/jmjoy/o"
)

func TestHashset(t *testing.T) {
	testdata := [...]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	hashset := o.NewHashSet()
	for _, v := range testdata {
		hashset.Add(v)
	}
	if len(testdata) != hashset.Size() {
		t.Fatal("HashSet Size error")
	}
	hashset.Remove(5)
	hashset.Remove(6)
	hashset.Remove(7)
	if len(testdata)-3 != hashset.Size() {
		t.Fatal("HashSet Size error")
	}

	if !hashset.Has(1) || hashset.Has(6) {
		t.Fatal("HashSet Has error")
	}

	testdata2 := [...]int{
		1, 4, 9, 16, 64, 81,
	}
	data := make([]int, 0, 6)
	hashset.Walk(func(value interface{}) error {
		intValue := value.(int)
		data = append(data, intValue*intValue)
		return nil
	})
	sort.Ints(data)
	for i := range testdata2 {
		if data[i] != testdata2[i] {
			t.Fatal("HashSet Walk error")
		}
	}

}
