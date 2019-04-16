## 类型
```
type byte = uint8

type rune = int32
```
可以看到二者类型并不一样。

golang中string底层是通过byte数组实现的。

```
package main

import (
	//"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	str := "hello world 你好"
	byteStr := []byte(str)
	runeStr := []rune(str)

	spew.Dump(str)
	spew.Dump(byteStr)
	spew.Dump(runeStr)
}
```
输出的结果如下：
```
(string) (len=18) "hello world 你好"
([]uint8) (len=18 cap=32) {
 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 20 e4 bd a0 e5  |hello world ....|
 00000010  a5 bd                                             |..|
}
([]int32) (len=14 cap=16) {
 (int32) 104,
 (int32) 101,
 (int32) 108,
 (int32) 108,
 (int32) 111,
 (int32) 32,
 (int32) 119,
 (int32) 111,
 (int32) 114,
 (int32) 108,
 (int32) 100,
 (int32) 32,
 (int32) 20320,
 (int32) 22909
}
```
可以看到string和[]byte输出了长度18， 而[]rune输出的长度为14。18是底层数组的长度，因为string的底层是[]string，而14是字符串真实长度。

中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。所以如果我们需要字符串的真实长度，可以将string转换为[]rune
