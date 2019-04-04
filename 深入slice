## slice数据结构
首先来看一下slice的数据结构
```
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
- array：指向所引用的数组指针
- len: 切片长度， 当前引用切片的元素个数
- cap: 切片容量（底层数组的元素总数）

cap始终是大于等于len的。

注意：

由于底层数组是可以被多个slice同时指向的，对一个slice的元素进行操作是有可能影响到其他slice的。

```
func main() {
	arrA := [3]int{1, 2, 3}
	fmt.Printf("arrA the address: %p\n", &arrA)

	arrB := arrA[:]
	fmt.Printf("arrB the address: %p\n", &arrB)

	arrC := arrA[:]
	arrC[2] = 4

	fmt.Println(arrA, arrB, arrC)
}
```
输出如下：
```
arrA the address: 0xc000016340
arrB the address: 0xc00000a060
[1 2 4] [1 2 4] [1 2 4]
```

可以看到通过修改了arrC[2]的值，arrA, arrB的值都修改了。

## slice创建

方式| 实例
---|---
直接声明 | var slice []int
new | slice := *new([]int)
字面量| slice := []int{1,2,3,4,5}
make| slice := make([]int, 3, 5)
从切片或数组截取| slice := arr[:]

切片操作并不复制切片指向的元素。它创建一个新的切片并复用原来切片的底层数组。 这使得切片操作和数组索引一样高效。因此，通过一个新切片修改元素会影响到原始切片的对应元素。

通过截取其他数组或者切片来创建切片，需要注意以下几点：
1. slice := arr[a:b]

这时slice的len和cap如下：
```
len = b - a
cap = len(arr) - a
```

2. slice := arr[a:b:c]

这时slice的len和cap如下：
```
len = b - a 
cap = c - a
```

3. 之前说过，多个slice指向同一个arr, 修
改其中一个slice会影响到其他的slice和原始arr。
```
arr := []int{1, 2, 3, 4, 5, 6}
	s1 := arr[2:4]
	s2 := arr[3:4:5]

	fmt.Println(arr, s1, s2, cap(s1), cap(s2))

	s2 = append(s2, 100)
	fmt.Println(arr, s1, s2, cap(s1))

	s2 = append(s2, 200)
	fmt.Println(arr, s1, s2, cap(s1))
```
输出的结果如下：
```
[1 2 3 4 5 6] [3 4] [4] 4 2
[1 2 3 4 100 6] [3 4] [4 100] 4
[1 2 3 4 100 6] [3 4] [4 100 200] 4
```
可以看到当执行s2 = append(s2, 200)这句话的时候，s2的cap不够了，需要扩容，重新开辟了一个地址，就不会影响到原有的底层数组了。

这里需要探讨以下append的情况。在进行append的时候，如果slice的容量够，那么就会修改原有底层数组， 如果容量不够， slice就得扩容，这样就不会产生影响。

```
s := []int{1, 2, 3, 4, 5, 6}
s1 := s[2:4]
s2 := s[3:6]

// 这时s1 = [3,4],  s2 = [4,5,6]
```

### nil切片和空切片
```
// nil切片
var nilSlice []int

// 空切片
nullSlice := []int{}
```

nil切片和空切片，它们的长度和容量都是0，但是它们指向底层数组的指针不一样，nil切片意味着指向底层数组的指针为nil，而空切片对应的指针是个地址。

可以理解为， nil切片并没有进行地址分配，而空切片已经分配了。

## append
slice一个特性就是可以通过append操作来进行扩容。

扩容的时候会调用growslice函数
```
func growslice(et *_type, old slice, cap int) slice {
    ...
    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        if old.len < 1024 {
            newcap = doublecap
        } else {
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            ...
        }
    }
    ...
}
```
扩容策略：
1. 如果新申请的容量大于原来容量的2倍，那么最后的容量就是新申请的容量。
2. 如果旧切片的长度小于1024， 那么最后的容量就是原来容量的2倍
3. 如果旧切片的长度大于1024， 那么最后的容量增长因子就是原来的1/4

