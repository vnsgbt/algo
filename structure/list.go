package list

// List //
type List struct {
	First *Item
	Last  *Item
}

// Item //
type Item struct {
	Content  interface{}
	Previous *Item
	Next     *Item
}

// NewItem //
func NewItem(value interface{}) *Item {
	return &Item{Content: value}
}

// NewFirst //
func (l *List) NewFirst(value interface{}) {
	item := NewItem(value)
	if l.Size() == 0 {
		l.First = item
		l.Last = l.First
	} else {
		item.Next = l.First
		l.First.Previous = item
		l.First = item
	}
}

// Size //
func (l *List) Size() int {
	if l.First == nil {
		return 0
	}
	return l.First.Count()
}

// Count //
func (i *Item) Count() int {
	if i.Next == nil {
		return 1
	}
	return 1 + i.Next.Count()
}
