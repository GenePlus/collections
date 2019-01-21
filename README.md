# 📂 collctions

> Golang 实现的 collections 模块，灵感来自 [Python queue](https://docs.python.org/3/library/queue.html) 和 [Python collections](https://docs.python.org/3/library/collections.html)

[![Build Status](https://travis-ci.org/chenjiandongx/collections.svg?branch=master)](https://travis-ci.org/chenjiandongx/collections) [![Go Report Card](https://goreportcard.com/badge/github.com/chenjiandongx/collections)](https://goreportcard.com/report/github.com/chenjiandongx/collections) [![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/chenjiandongx/collections?status.svg)](https://godoc.org/github.com/chenjiandongx/collections)

## 📚 目录
* [Queue - 先进先出队列](#Queue)
* [LifoQueue - 后进先出队列](#LifoQueue)
* [PriorityQueue - 优先队列](#PriorityQueue)
* [Deque - 双端队列](#Deque)
* [OrderedMap - 有序 Map](#OrderedMap)
* [Counter - 计数器](#Counter)
* [Sort - 排序](#Sort)

### 🔰 安装&引用
```bash
$ go get github.com/chenjiandongx/collections

import "github.com/chenjiandongx/collections"
```

### 📦 Collections
### Queue
> 先进先出队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v interface{})          // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000

q := collections.NewQueue()
var item interface{}
var ok bool
for i := 0; i < nums; i++ {
    q.Put(i)
}
for i := 0; i < nums; i++ {
    if item, ok = q.Get(); ok {
        fmt.Println(item.(int))
    }
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### LifoQueue
> 后进先出队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v interface{})          // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000

q := collections.NewLifoQueue()
var item interface{}
var ok bool
for i := 0; i < nums; i++ {
    q.Put(i)
}
for i := nums-1; i >=0; i-- {
    if item, ok = q.Get(); ok {
        fmt.Println(item.(int))
    }
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### PriorityQueue
> 优先队列（线程安全）

📝 方法集
```shell
Get()(interface{}, bool)    // 出队
Put(v *PqNode)              // 入队
Qsize() int                 // 返回队列长度
IsEmpty() bool              // 判断队列是否为空

// 优先队列节点
type PqNode struct {
    Value           string
    Priority, index int
}
```

✏️ 示例
```go
var nums = 1000

q := collections.NewPriorityQueue()

for i := 0; i < nums; i++ {
    r := rand.Int()
    q.Put(&collections.PqNode{Value: string(r), Priority: rand.Int()})
}

for i := 0; i < nums/2; i++ {
    item1, _ := q.Get()
    item2, _ := q.Get()
    fmt.Println(item1.(*collections.PqNode).Priority > item2.(*collections.PqNode).Priority)
}

fmt.Println(q.IsEmpty())
fmt.Println(q.Qsize())
```

### Deque
> 双端队列（线程安全）

📝 方法集
```shell
GetLeft()(interface{}, bool)        // 左边出队
GetRight()(interface{}, bool)       // 右边出队
PutLeft(v interface{})              // 左边入队
PutRight(v interface{})             // 右边入队
Qsize() int                         // 返回队列长度
IsEmpty() bool                      // 判断队列是否为空
```

✏️ 示例
```go
var nums = 1000
q := collections.NewDeque()

var item interface{}
var ok bool

for i := 0; i < nums; i++ {
    q.PutLeft(i)
}
fmt.Println(q.Qsize())

for i := nums - 1; i >= 0; i-- {
    q.PutRight(i)
}
fmt.Println(q.Qsize())

for i := 0; i < nums; i++ {
    item, ok = q.GetRight()
    fmt.Println(item, ok)
}
for i := nums - 1; i >= 0; i-- {
    item, ok = q.GetLeft()
    fmt.Println(item, ok)
}

item, ok = q.GetLeft()
fmt.Println(item, ok)

item, ok = q.GetRight()
fmt.Println(item, ok)
```

### OrderedMap
> 有序 Map，接口设计参考 [cevaris/ordered_map](https://github.com/cevaris/ordered_map)

📝 方法集
```shell
Set(key, value interface{})                 // 新增键值对
Get(key interface{}) (interface{}, bool)    // 取值
Delete(key interface{}) bool                // 删除键
Iter() (interface{}, interface{}, bool)     // 遍历
Len() int                                   // 键值对数量
// 指针回退到 Head，遍历时 current 指针会向后移动 BackToHead 使其移动到头指针，以便下一次从头遍历
BackToHead()                               
```

✏️ 示例
```go
maxNum := 100
om := collections.NewOrderedMap()
for i := 0; i < maxNum; i++ {
    om.Set(i, i+1)
}

fmt.Println(om.Len())
om.Delete(0)
fmt.Println(om.Len())

for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
    fmt.Println(k, v)
}

om.BackToHead()
for k, v, ok := om.Iter(); ok; k, v, ok = om.Iter() {
    fmt.Println(k, v)
}
```

📣 讨论

有序 Map 在 Golang 中应该是一个十分常见的需求，Map 最大的优势就是它的查找性能，**理论上** Map 查找的时间复杂度是常数级的。但实际情况如何，我们可以通过 benchmark 来验证。在 [Go Maps Don’t Appear to be O(1)](https://medium.com/@ConnorPeet/go-maps-are-not-o-1-91c1e61110bf) 这篇文章中，作者测试了 Golang Map 查找的实际性能，不过作者是基于 Go1.4 的，版本有点旧了。下面是我修改了作者的测试案例后在 Go1.10 下跑出来的结果。

![](https://user-images.githubusercontent.com/19553554/51075377-83f8cd80-16c5-11e9-9973-4904a4661aeb.png)

上图是使用 [go-echarts](https://github.com/chenjiandongx/go-echarts) 绘制的。测试是通过与二分查找来对比的，二分查找的时间复杂度为 **O(log2n)**。很明显，在 10e5 数量级下两者的性能差别还不是特别大，主要差距是在 10e6 后体现的。结论：Map 的性能优于 **O(log2n)**，但不是常数级。

**collections.OrderdMap 🆚 cevaris/ordered_map**

本来我一直使用的是 [cevaris/ordered_map](https://github.com/cevaris/ordered_map)，后来自己重新实现了一个。实现完就与它进行了性能测试对比。它是基于两个 Map 实现的，而我是使用的 Map+LinkedList，LinkedList 在删除和插入操作上的时间复杂度都是 **O(1)**，用其来存储 Map key 的顺序是一个很好的选择。

同样的测试代码，BenchMark 结果如下
```shell
goos: windows
goarch: amd64
pkg: github.com/chenjiandongx/collections
BenchmarkCollectionsSet-8        2000000               689 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisSet-8            1000000              1212 ns/op             334 B/op          3 allocs/op
BenchmarkCollectionsGet-8        2000000               823 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisGet-8            1000000              1281 ns/op             334 B/op          3 allocs/op
BenchmarkCollectionsIter-8       2000000               670 ns/op             187 B/op          3 allocs/op
BenchmarkCevarisIter-8           1000000              1341 ns/op             366 B/op          4 allocs/op
```
**collections.OrderedMap Win 🖖 性能+内存占用全部占优 🚀**

### Counter
> 计数器

📝 方法集
```shell
// key-value item
type Item struct {
    k interface{}
    v int
}

Add(keys ...interface{})            // 新增 item
Get(key interface{}) int            // 获取 key 计数
GetAll() []Item                     // 获取全部 key 计数
Top(n int) []Item                   // 获取前 key 计数
Delete(key interface{}) bool        // 删除 key，成功返回 true，key 不存在返回 false
Len() int                           // key 数量
```

✏️ 示例
```go
c := collections.NewCounter()
c.Add("a", "b", "c", "d", "a", "c")
fmt.Println(c.Get("A"))
fmt.Println(c.Get("a"))
fmt.Println(c.Get("b"))
fmt.Println(c.Top(2))
fmt.Println(c.Len())
fmt.Println(c.All())
c.Delete("a")
```

### Sort

📝 方法集
```shell
BubbleSort()    // 冒泡排序
InsertSort()    // 插入排序
QuickSort()     // 快速排序
ShellSort()     // 希尔排序
HeapSort()      // 堆排序
MergeSort()     // 归并排序
```

### 📃 License
MIT [©chenjiandongx](http://github.com/chenjiandongx)
