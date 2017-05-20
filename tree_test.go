package tree

import (
	"math/rand"
	"sort"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestValue float64

func (t TestValue) Index() float64 {
	return float64(t)
}

type TreeSuite struct {
	t *Tree
	r *rand.Rand
}

var _ = Suite(&TreeSuite{})

func (s *TreeSuite) SetUpTest(c *C) {
	var f Node = TestValue(0.0)
	s.t = New(&f)
	s.r = rand.New(rand.NewSource(88))
}

func (s *TreeSuite) TestNew(c *C) {
	c.Check(s.t.Left, IsNil)
	c.Check(s.t.Right, IsNil)
	c.Check((*s.t.Value).Index(), Equals, 0.0)
}

func (s *TreeSuite) TestInsert(c *C) {
	values := make([]float64, 11)
	values[0] = 0.0 // The test tree is initialized with 0.

	// Generate values and insert them into the tree.
	for i := 1; i < cap(values); i++ {
		var f Node
		values[i] = s.r.Float64()
		f = TestValue(values[i])
		s.t.Insert(&f)
	}
	// Sort the values because we expect the in-order traversal
	// to visit the nodes in order.
	sort.Float64s(values)

	// Anonymous function to connect the order the nodes
	// are visited in.
	results := make([]float64, 0, 11)
	collect := func(t *Tree) *Tree {
		results = results[0 : len(results)+1]
		results[len(results)-1] = (*t.Value).Index()
		return t
	}

	// Traverse the tree to collect results
	s.t.Traverse(collect)

	// Verify results
	c.Check(results, DeepEquals, values)
}
