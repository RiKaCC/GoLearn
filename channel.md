# channel
## 基本知识
channel跟Map一样，是一个使用make创建的数据结构的引用。

当复制或者作为参数传递到一个函数时，复制的是引用。

通道的零值是nil，通道可以使用==来进行比较。

## 创建一个channle

```
// 创建一个无缓冲的通道
ch := make(chan int) // ch的类型是 chan int
// 创建一个有缓冲通道
ch := make(chan int, 3) // 容量为3的缓冲通道
```
### 无缓冲通道
无缓冲通道上的发送操作将会被阻塞，直到另一个goroutine在对应的通道上执行接收操作，这时值传送完成。两个goroutine都能继续执行。

因此无缓冲通道也叫做同步通道，当一个同步通道。

#### 一个简单示例
使用三个goroutine来计算平方数。

一个goroutine产生数，一个goroutine计算平方，一个goroutine来输出平方数。

```
package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			a <- i
		}
		close(a)
	}()

	go func() {
		for x := range a {
			b <- x * x
		}
		close(b)
	}()

	for x := range b {
		fmt.Println(x)
	}
}
```
需要注意的地方就是，通道的close。

如果不执行close(a)，就会报错`fatal error: all goroutines are asleep - deadlock!`.

通过报错说明，我们能知道产生了死锁，确切的说，是所有的goroutine都sleep了。

为什么会出现这种情况？

如果不close(a),当i=10的时候，就不会往channle a里面写了，但是并没有关闭这个通道。在计算平方的goroutine并不知道，他就阻塞了，一直等着通道a有值。也不会向通道b写值。而输出的goroutine一直在等待通道b, 这样就产生了死锁，导致了这个报错。

### 关闭通道
如果发送方知道已经没有更多的数据需要发送了，告诉接收的goroutine可以停止等待是非常有用的。这时候可以通过调用close函数来关闭通道。

在关闭通道之后，如果再往这个通道里写数据的话，会导致程序崩溃。

在golang中，并没有一个很直接的方式来判断通道是否被关闭了，我们可以采用range循环的方式来迭代通道。

关于通道的回收，通道的回收是GC的时候根据它是否可以访问来决定是否回收，并不是根据它是否关闭。
