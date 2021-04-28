package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCounter(t *testing.T) {
	c := NewCounter()
	assert.Equal(t, c.Len(), 0)
	assert.Equal(t, len(c.GetAll()), 0)
	assert.Equal(t, c.Get("anything"), 0)
}

func TestCounter(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	assert.Equal(t, c.Top(2), []Item{{"c", 3}, {"a", 2}})
	assert.Equal(t, c.Get("A"), 0)
	assert.Equal(t, c.Get("a"), 2)
	assert.Equal(t, c.Get("b"), 1)
	assert.Equal(t, c.Len(), 4)
	assert.Equal(t, len(c.Top(10)), c.Len())
	assert.Equal(t, len(c.Top(-10)), c.Len())
	c.Delete("a")
	assert.Equal(t, c.Get("a"), 0)
	assert.Equal(t, c.Len(), 3)
}

func TestGetKeys(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	cString := c.GetKeysString()
	assert.Equal(t, cString[0], "c")
	assert.Equal(t, len(cString), 4)
}

func TestCounterAddConst(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	assert.Equal(t, c.Get("a"), 2)
	assert.Equal(t, c.Len(), 4)
	c.AddConst("a", 3)
	assert.Equal(t, c.Get("a"), 5)
	c.AddConst("e", 3)
	assert.Equal(t, c.Get("e"), 3)
	assert.Equal(t, c.Len(), 5)
}

// func TestCounterMapAdd(t *testing.T) {
// 	mc := NewMapCounter()
// 	x := NewCounter()
// 	x.Add("a", "b", "c", "d", "a", "c", "c")
// 	mc["X"] = x
// 	mc["X"].Add("b")
// 	assert.Equal(t, mc["X"].Get("a"), 2)
// 	assert.Equal(t, mc["X"].Get("b"), 2)
// 	assert.Equal(t, mc["X"].Len(), 4)
// 	mc.MapCounterAdd("X", "a")
// 	mc.MapCounterAdd("X", "e")
// 	assert.Equal(t, mc["X"].Get("a"), 3)
// 	assert.Equal(t, mc["X"].Get("e"), 1)
// 	assert.Equal(t, mc["X"].Len(), 5)
// }

func TestCounterMapAdd(t *testing.T) {
	mc := NewMapCounter()
	x := NewCounter()
	x.Add("a", "b", "c", "d", "a", "c", "c")
	mc[1] = x
	mc[1].Add("b")
	assert.Equal(t, mc[1].Get("a"), 2)
	assert.Equal(t, mc[1].Get("b"), 2)
	assert.Equal(t, mc[1].Len(), 4)
	mc.MapCounterAdd(1, "a")
	mc.MapCounterAdd(1, "e")
	assert.Equal(t, mc[1].Get("a"), 3)
	assert.Equal(t, mc[1].Get("e"), 1)
	assert.Equal(t, mc[1].Len(), 5)
}

func TestCounterMapAddConst(t *testing.T) {
	mc := NewMapCounter()
	mc.MapCounterAddConst("X", "a", 3)
	mc.MapCounterAddConst("X", "b", 2)
	assert.Equal(t, mc["X"].Get("a"), 3)
	assert.Equal(t, mc["X"].Get("b"), 2)
	assert.Equal(t, mc["X"].Len(), 2)
}

func TestCounterDel(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	assert.Equal(t, c.Get("a"), 2)
	assert.Equal(t, c.Len(), 4)
	c.Del("a", "c", "d")
	assert.Equal(t, c.Get("a"), 1)
	assert.Equal(t, c.Get("c"), 2)
	assert.Equal(t, c.Len(), 3)
	assert.Equal(t, c.Get("d"), 0)
	c.Delete("a")
	assert.Equal(t, c.Get("a"), 0)
	assert.Equal(t, c.Len(), 2)
}

func TestCounterDelConst(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	assert.Equal(t, c.Get("a"), 2)
	assert.Equal(t, c.Len(), 4)
	c.DelConst("c", 2)
	assert.Equal(t, c.Get("c"), 1)
	c.DelConst("a", 2)
	assert.Equal(t, c.Get("a"), 0)
	assert.Equal(t, c.Len(), 3)
}

func TestCounterMapDel(t *testing.T) {
	mc := NewMapCounter()
	x := NewCounter()
	x.Add("a", "b", "c", "d", "a", "c", "c")
	mc["X"] = x
	mc["X"].Add("b")
	mc.MapCounterDel("X", "c")
	mc.MapCounterDel("X", "b")
	assert.Equal(t, mc["X"].Get("c"), 2)
	assert.Equal(t, mc["X"].Get("b"), 1)
	assert.Equal(t, mc["X"].Len(), 4)
	mc.MapCounterDel("X", "d")
	mc.MapCounterDel("X", "e")
	assert.Equal(t, mc["X"].Len(), 3)
}

func TestCounterMapDelConst(t *testing.T) {
	mc := NewMapCounter()
	x := NewCounter()
	x.Add("a", "b", "c", "d", "a", "c", "c")
	mc["X"] = x
	assert.Equal(t, mc["X"].Len(), 4)
	mc.MapCounterDelConst("X", "c", 2)
	mc.MapCounterDelConst("X", "b", 1)
	assert.Equal(t, mc["X"].Get("c"), 1)
	assert.Equal(t, mc["X"].Get("b"), 0)
	assert.Equal(t, mc["X"].Len(), 3)
	mc.MapCounterDelConst("X", "d", 1)
	mc.MapCounterDelConst("X", "e", 2)
	assert.Equal(t, mc["X"].Len(), 2)
}

func TestCounterSumAll(t *testing.T) {
	c := NewCounter()
	c.Add("a", "b", "c", "d", "a", "c", "c")
	assert.Equal(t, c.SumAll(), 7)
	assert.Equal(t, c.Len(), 4)
}
