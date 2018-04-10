package list

import "testing"

func TestList(t *testing.T) {

	l := new(List)
	l.NewFirst(NewItem(1))
	l.NewFirst(NewItem(2))
	l.NewFirst(NewItem(3))

	if l.Size() != 3 {
		t.Error()
	}
}
