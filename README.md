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
* [AVLTree - AVL 树](#AVLTree)
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

### AVLTree
> AVL 二叉自平衡查找树

📝 方法集
```shell
NewAVLTree() *AVLTree       // 生成 AVL 树
Insert(v int)               // 插入节点
Search(v int) bool          // 搜索节点
Delete(v int) bool          // 删除节点
GetMaxValue() int           // 获取所有节点中的最大值
GetMinValue() int           // 获取所有节点中的最小值
AllValues() []int           // 返回排序后所有值
```

✏️ 示例
```go
var maxNum = 100

tree := NewAVLTree()
for i := 0; i < maxNum; i++ {
    tree.Insert(i)
    tree.Insert(maxNum + i)
}
fmt.Println(len(tree.AllValues()))
fmt.Println(tree.GetMaxValue())
fmt.Println(tree.GetMinValue())
fmt.Println(tree.Search(50))
fmt.Println(tree.Search(100))
fmt.Println(tree.Search(-10))
fmt.Println(tree.Delete(-10))
fmt.Println(tree.Delete(10))
```

📣 讨论

AVL 树是自平衡树的一种，其通过左旋和右旋来调整自身的平衡性，使其左右子树的高度差最大不超过 1。AVL 在插入、查找、删除的平时时间复杂度都是 O(logn)，在基本的 BST（二叉查找树）中，理想情况的效率也是为 O(logn)，但由于操作的性能其实是依赖于树的高度，而 BST 最坏的情况会导致树退化成链表，此时时间复杂度就变为 O(n)，为了解决这个问题，自平衡二叉树应运而生。

AVL 的主要精髓在于`旋转`，旋转分为 4 种情况，左旋，左旋+右旋，右旋，右旋+左旋。调整树结构后需要重新计算树高。

**左子树左节点失衡**
> 左左情况 直接右旋
```shell
    x                
  x        => 右旋         x
x                       x    x
```

**左子树右节点失衡**
> 左右情况 先左旋后右旋
```shell
  x                        x     
x         => 左旋         x       => 右旋        x
  x                     x                     x    x
```

**右子树右节点失衡**
> 右右情况 直接左旋
```shell
x                
  x       => 左旋          x
    x                   x    x
```

**右子树左节点失衡**
> 右左情况 先右旋后左旋
```shell
x                      x     
  x       => 右旋        x       => 左旋        x
x                          x                 x    x
```

AVL 主要的性能消耗主要在插入，因为其需要通过旋转来维护树的平衡，但如果使用场景是经常需要排序和查找数据的话，AVL 还是可以展现其良好的性能的。

### Sort

📝 方法集
```shell
BubbleSort()    // 冒泡排序
InsertionSort()    // 插入排序
QuickSort()     // 快速排序
ShellSort()     // 希尔排序
HeapSort()      // 堆排序
MergeSort()     // 归并排序
```

✏️ 示例
```go
var maxCnt = 10e4

func yieldRandomArray() []int {
    res := make([]int, maxCnt)
    for i := 0; i < maxCnt; i++ {
        res[i] = rand.Int()
    }
    return res
}

BubbleSort(yieldRandomArray())
InsertionSort(yieldRandomArray())
QuickSort(yieldRandomArray())
ShellSort(yieldRandomArray())
HeapSort(yieldRandomArray())
MergeSort(yieldRandomArray())
```

📣 讨论

**排序算法时间复杂度比较**

| 排序算法 |  是否稳定  |  平均    |   最好  |    最差   |   动画演示  |
| -------- | --------- |----------| --------| -------- | ----------- |
| BubbleSort | 是 | O(n^2) |  O(n) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/3/37/Bubble_sort_animation.gif) |
| InsertionSort | 是 | O(n^2) |  O(n) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif) |
| QuickSort | 否 | O(nlogn) | O(nlogn) |  O(n^2) | ![](https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif) |
| ShellSort | 否 |O(nlogn) |  O(n) | O(n^2)  | ![](https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif) |
| HeapSort | 否 | O(nlogn) |  O(nlogn) | O(nlogn) | ![](https://upload.wikimedia.org/wikipedia/commons/1/1b/Sorting_heapsort_anim.gif) |
| MergeSort | 是 | O(nlogn) |  O(nlogn) | O(nlogn) | ![](https://upload.wikimedia.org/wikipedia/commons/c/c5/Merge_sort_animation2.gif) |

具体运行时间如何呢，可以通过 benchmark 来测试一下
```go
// 生成指定长度的随机整数数组
var maxCnt int = 10e4

func yieldRandomArray(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = rand.Int()
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8                  1        17361549400 ns/op
BenchmarkInsertionSort-8               1        1934826900 ns/op
BenchmarkQuickSort-8                 100          10651807 ns/op
BenchmarkShellSort-8                 100          16476199 ns/op
BenchmarkHeapSort-8                  100          14231607 ns/op
BenchmarkMergeSort-8                 100          14840583 ns/op
```

换两种极端的数据分布方式
```go
// 升序
func yieldArrayAsce(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = i
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8               5000            266690 ns/op
BenchmarkInsertionSort-8           10000            213429 ns/op
BenchmarkStdSort-8                   200           6901535 ns/op
BenchmarkQuickSort-8                   1        3291222900 ns/op
BenchmarkShellSort-8                1000           1716406 ns/op
BenchmarkHeapSort-8                  200           6806788 ns/op
BenchmarkMergeSort-8                 300           4677485 ns/op
```

```go
// 降序
func yieldArrayDesc(cnt int) []int {
    res := make([]int, cnt)
    for i := 0; i < cnt; i++ {
        res[i] = cnt-i
    }
    return res
}
```

运行结果
```shell
BenchmarkBubbleSort-8                  1        6710048800 ns/op
BenchmarkInsertionSort-8               1        3881599100 ns/op
BenchmarkQuickSort-8                   1        3373971200 ns/op
BenchmarkShellSort-8                 500           2876371 ns/op
BenchmarkHeapSort-8                  200           7081150 ns/op
BenchmarkMergeSort-8                 300           4448222 ns/op
```

// TODO: 睡觉先

### 📃 License
MIT [©chenjiandongx](http://github.com/chenjiandongx)
