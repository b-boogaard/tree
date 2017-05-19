package tree

// Tree is a recursive data structure representing a binary tree.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New creates a fresh zero defaulted Tree.
func New(v int) *Tree {
	return &Tree{nil, v, nil}
}

// Insert adds value v into the proper position within
// the tree.
func (t *Tree) Insert(v int) *Tree {
	if t == nil {
		return New(v)
	}

	if v < t.Value {
		t.Left = t.Left.Insert(v)
	} else {
		t.Right = t.Right.Insert(v)
	}

	return t
}

func minValue(t *Tree) int {
	min := t.Value

	for t.Left != nil {
		min = t.Left.Value
		t = t.Left
	}

	return min
}

// Delete removes the node with value v from the tree.
// When successful the deleted node will be returned.
func (t *Tree) Delete(v int) *Tree {
	switch {
	case t == nil:
		return t
	case v < t.Value:
		t.Left = t.Left.Delete(v)
	case v > t.Value:
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
func (t *Tree) Find(v int) *Tree {
	if t == nil || t.Value == v {
		return t
	}

	if v < t.Value {
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
