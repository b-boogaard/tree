package tree

// Node is an interface that allows
// for many different types to be used
// within a tree so long as they define
// some sort of index that can be represented
// as a float32.
type Node interface {
	Index() float64
}

// Tree is a recursive data structure representing a binary tree.
type Tree struct {
	Left  *Tree
	Value *Node
	Right *Tree
}

// New creates a fresh zero defaulted Tree.
func New(v *Node) *Tree {
	return &Tree{nil, v, nil}
}

// Insert adds value v into the proper position within
// the tree.
func (t *Tree) Insert(v *Node) *Tree {
	if t == nil {
		return New(v)
	}

	if (*v).Index() < (*t.Value).Index() {
		t.Left = t.Left.Insert(v)
	} else {
		t.Right = t.Right.Insert(v)
	}

	return t
}

func minValue(t *Tree) *Node {
	min := t.Value

	for t.Left != nil {
		min = t.Left.Value
		t = t.Left
	}

	return min
}

// Delete removes the node with value v from the tree.
// When successful the deleted node will be returned.
func (t *Tree) Delete(v *Node) *Tree {
	switch {
	case t == nil:
		return t
	case (*v).Index() < (*t.Value).Index():
		t.Left = t.Left.Delete(v)
	case (*v).Index() > (*t.Value).Index():
		t.Right = t.Right.Delete(v)
	default:
		if t.Left == nil {
			return t.Right
		}

		if t.Right == nil {
			return t.Left
		}

		t.Value = minValue(t.Right)
		t.Right = t.Right.Delete(t.Value)
	}

	return t
}

// Find locates value v and returns the Tree node
// containing it.
func (t *Tree) Find(v *Node) *Tree {
	if t == nil || (*t.Value).Index() == (*v).Index() {
		return t
	}

	if (*v).Index() < (*t.Value).Index() {
		return t.Left.Find(v)
	}

	return t.Right.Find(v)
}

// Traverse does an in-order traversal of the tree.
func (t *Tree) Traverse(fn func(*Tree) *Tree) {
	if t == nil {
		return
	}

	t.Left.Traverse(fn)
	fn(t)
	t.Right.Traverse(fn)

	return
}
