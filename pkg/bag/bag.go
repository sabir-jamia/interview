package bag

// Bag is a container
type Bag struct {
	first *node
	n int
}

type node struct {
	item int
	next *node
}

// Init initializes or clears Bag b.
func(b *Bag) Init() *Bag {
	b.first = nil
	b.n = 0
	return b
}

// New returns an initialized bag.
func New() *Bag {
	return new(Bag).Init() 
}

// Add adds an item to the bag
func (b *Bag) Add(item int) {
	oldFirst := b.first
	b.first = new(node)
	b.first.item = item
	b.first.next = oldFirst
	b.n++
}

// IsEmpty checks if bag is empty
func (b *Bag) IsEmpty() bool{
	return b.first == nil
}

// Size returns the size of bag
func (b *Bag) Size() int {
	return b.n
}

// Iterate returns a slice of items
func (b *Bag) Iterate() []int {
	var items []int
	for itemNode := b.first; itemNode != nil; itemNode = itemNode.next {
		items = append(items, itemNode.item)
	}
	return items
}