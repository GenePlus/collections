// update from github.com/chenjiandongx/collections
package collections

import (
	"sort"
)

type Counter struct {
	kv map[interface{}]int
}

type Item struct {
	k interface{}
	v int
}

// GetK, do not change k, v as K, V due to being used
func (p *Item) GetK() interface{} {
	return p.k
}

// GetV
func (p *Item) GetV() int {
	return p.v
}

func NewCounter() *Counter {
	return &Counter{make(map[interface{}]int)}
}

func (c *Counter) Add(keys ...interface{}) {
	for i := 0; i < len(keys); i++ {
		c.kv[keys[i]]++
	}
}

// AddConst is add num times for the one key.
func (c *Counter) AddConst(s1 interface{}, num int) {
	for i := 0; i < num; i++ {
		c.Add(s1)
	}
}

// GetValuesInt
func (c *Counter) GetValuesInt() []int {
	rValues := make([]int, 0, c.Len())
	for _, itm := range c.GetAll() {
		rValues = append(rValues, itm.v)
	}
	return rValues
}

// GetKeysString
func (c *Counter) GetKeysString() []string {
	rKeys := make([]string, 0, c.Len())
	for _, itm := range c.GetAll() {
		rKeys = append(rKeys, itm.k.(string))
	}
	return rKeys
}

// GetKeysInt
func (c *Counter) GetKeysInt() []int {
	rKeys := make([]int, 0, c.Len())
	for _, itm := range c.GetAll() {
		rKeys = append(rKeys, itm.k.(int))
	}
	return rKeys
}

// GetKeysFloat64
func (c *Counter) GetKeysFloat64() []float64 {
	rKeys := make([]float64, 0, c.Len())
	for _, itm := range c.GetAll() {
		rKeys = append(rKeys, itm.k.(float64))
	}
	return rKeys
}

// Del is value num -1, differ with Delete which remove the key-value.
func (c *Counter) Del(keys ...interface{}) {
	for i := 0; i < len(keys); i++ {
		if num, ok := c.kv[keys[i]]; ok {
			if num > 1 {
				c.kv[keys[i]]--
			} else if num == 1 {
				delete(c.kv, keys[i])
			}
		}

	}
}

// DelConst is del num times for the one key.
func (c *Counter) DelConst(s1 interface{}, n int) {
	//num := int(math.Abs(float64(n)))
	//if c.Get(s1) >= num {
	for i := 0; i < n; i++ {
		c.Del(s1)
	}
	//}
}

// MapCounter is a map with value as Counter.
type MapCounter map[interface{}]*Counter

// NewMapCounter need to init the inner 2nd map.
func NewMapCounter() MapCounter {
	mc := make(map[interface{}]*Counter)
	for i, _ := range mc {
		mc[i] = NewCounter()
	}
	return mc
}

func (mc MapCounter) MapCounterAdd(s1 interface{}, s2 interface{}) {
	if _, ok := mc[s1]; ok {
		mc[s1].Add(s2)
	} else {
		c := NewCounter()
		c.Add(s2)
		mc[s1] = c
	}
}
func (mc MapCounter) MapCounterAddConst(s1 interface{}, s2 interface{}, num int) {
	if _, ok := mc[s1]; ok {
		mc[s1].AddConst(s2, num)
	} else {
		c := NewCounter()
		c.AddConst(s2, num)
		mc[s1] = c
	}
}

func (mc MapCounter) MapCounterDel(s1 interface{}, s2 interface{}) {
	if _, ok := mc[s1]; ok {
		mc[s1].Del(s2)
	} else {
		c := NewCounter()
		c.Del(s2)
		mc[s1] = c
	}
}
func (mc MapCounter) MapCounterDelConst(s1 interface{}, s2 interface{}, num int) {
	if _, ok := mc[s1]; ok {
		mc[s1].DelConst(s2, num)
	} else {
		c := NewCounter()
		c.DelConst(s2, num)
		mc[s1] = c
	}
}

func (mc MapCounter) GetKeysString() []string {
	rKeys := make([]string, 0, len(mc))
	for key := range mc {
		rKeys = append(rKeys, key.(string))
	}
	return rKeys
}

func (mc MapCounter) GetKeysInt() []int {
	rKeys := make([]int, 0, len(mc))
	for key := range mc {
		rKeys = append(rKeys, key.(int))
	}
	return rKeys
}

func (c *Counter) Get(key interface{}) int {
	return c.kv[key]
}

func (c *Counter) GetAll() []Item {
	return c.sortMap()
}

func (c *Counter) Top(n int) []Item {
	sortItems := c.sortMap()
	if n > c.Len() || n < 0 {
		n = c.Len()
	}
	return sortItems[:n]
}

func (c *Counter) Delete(key interface{}) bool {
	if _, ok := c.kv[key]; ok {
		delete(c.kv, key)
		return true
	}
	return false
}

func (c *Counter) Len() int {
	return len(c.kv)
}

func (c *Counter) sortMap() []Item {
	var items []Item
	for k, v := range c.kv {
		items = append(items, Item{k, v})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].v > items[j].v
	})
	return items
}

// SumAll returns all Counter sum, differ with Len() which is deduped.
func (c *Counter) SumAll() int {
	var sum int
	for _, v := range c.kv {
		sum += v
	}
	return sum
}
