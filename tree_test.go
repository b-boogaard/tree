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
	i float64
	t *Tree
	r *rand.Rand
}

var _ = Suite(&TreeSuite{})

func initialize(r *rand.Rand) (*Tree, float64) {
	var initial float64 = r.Float64()
	var f Node = TestValue(initial)
	return New(f), initial
}

func populate(t *Tree, r *rand.Rand) []float64 {
	values := make([]float64, 11)

	if t == nil {
		t, values[0] = initialize(r)
	} else {
		values[0] = t.Value.Index()
	}

	for i := 1; i < cap(values); i++ {
		values[i] = r.Float64()
		var f Node = TestValue(values[i])
		t.Insert(f)
	}

	return values
}

func (s *TreeSuite) SetUpTest(c *C) {
	s.r = rand.New(rand.NewSource(88))
	s.t, s.i = initialize(s.r)
}

func (s *TreeSuite) TestNew(c *C) {
	c.Check(s.t.Left, IsNil)
	c.Check(s.t.Right, IsNil)
	c.Check(s.t.Value.Index(), Equals, s.i)
}

type max struct {
	value float64
}

func (s *TreeSuite) TestTraverse(c *C) {
	var maxValue float64

	traverseFunc := func(t *Tree) *Tree {
		if maxValue < t.Value.Index() {
			maxValue = t.Value.Index()
		}
		return t
	}

	values := populate(s.t, s.r)
	sort.Float64s(values)

	s.t.Traverse(traverseFunc)

	c.Check(maxValue, Equals, values[len(values)-1])
}

func (s *TreeSuite) TestInsert(c *C) {
	values := populate(s.t, s.r)

	// Sort the values because we expect the in-order traversal
	// to visit the nodes in order.
	sort.Float64s(values)

	// Anonymous function to connect the order the nodes
	// are visited in.
	results := make([]float64, 0, 11)
	collect := func(t *Tree) *Tree {
		results = results[0 : len(results)+1]
		results[len(results)-1] = t.Value.Index()
		return t
	}

	// Traverse the tree to collect results
	s.t.Traverse(collect)

	// Verify results
	c.Check(results, DeepEquals, values)
}

func (s *TreeSuite) TestFind(c *C) {
	values := populate(s.t, s.r)
	value := values[s.r.Intn(len(values))]
	var valueNode Node = TestValue(value)

	result := s.t.Find(valueNode)

	c.Check(result.Value.Index(), Equals, value)
}

func (s *TreeSuite) TestDelete(c *C) {
	values := populate(s.t, s.r)
	sort.Float64s(values)

	// validate all the values exist in the tree
	results := make([]float64, 0, 11)
	collect := func(t *Tree) *Tree {
		results = results[0 : len(results)+1]
		results[len(results)-1] = t.Value.Index()
		return t
	}

	s.t.Traverse(collect)

	c.Check(results, DeepEquals, values)

	// Select random node to delete
	value := values[s.r.Intn(len(values))]
	var valueNode Node = TestValue(value)

	// Remove node
	root := s.t.Delete(valueNode)

	// Validate root node is returned
	c.Check(root.Value.Index(), Equals, s.t.Value.Index())

	// filter the deleted value from the values slice
	filteredValues := values[:0]
	for _, x := range values {
		if x != value {
			filteredValues = append(filteredValues, x)
		}
	}

	c.Check(len(filteredValues), Equals, len(values)-1)

	// Collect all nodes from tree after delete
	secondResults := make([]float64, 0, 11)
	secondCollect := func(t *Tree) *Tree {
		secondResults = secondResults[0 : len(secondResults)+1]
		secondResults[len(secondResults)-1] = t.Value.Index()
		return t
	}

	s.t.Traverse(secondCollect)

	c.Check(secondResults, DeepEquals, filteredValues)
}
