package tree

import (
	"math/rand"
	"sort"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TreeSuite struct {
	t *Tree
	r *rand.Rand
}

var _ = Suite(&TreeSuite{})

func (s *TreeSuite) SetUpTest(c *C) {
	s.t = New(0)
	s.r = rand.New(rand.NewSource(88))
}

func (s *TreeSuite) TestNew(c *C) {
	c.Check(s.t.Left, IsNil)
	c.Check(s.t.Right, IsNil)
	c.Check(s.t.Value, Equals, 0)
}

func (s *TreeSuite) TestInsert(c *C) {
	values := make([]int, 11)
	values[0] = 0 // The test tree is initialized with 0.

	// Generate values and insert them into the tree.
	for i := 1; i < cap(values); i++ {
		values[i] = s.r.Int()
		s.t.Insert(values[i])
	}
	// Sort the values because we expect the in-order traversal
	// to visit the nodes in order.
	sort.Ints(values)

	// Anonymous function to connect the order the nodes
	// are visited in.
	results := make([]int, 0, 11)
	collect := func(t *Tree) *Tree {
		results = results[0 : len(results)+1]
		results[len(results)-1] = t.Value
		return t
	}

	// Traverse the tree to collect results
	s.t.Traverse(collect)

	// Verify results
	c.Check(results, DeepEquals, values)
}
