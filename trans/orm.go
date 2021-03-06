package trans


type RMNode interface {

	Parent() RMNode

	Child(index int) RMNode

	Children(start int, end int) []RMNode

	Size() int

	Value() interface{}

}