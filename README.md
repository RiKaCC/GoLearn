# Go-demo
## string package usage
#### string reverse
```golang

str := "reverse string"
ret := StringReverse(str) 
fmt.Println(ret)  //"gnirts esrever"
```
#### string join
```golang
//there two function for use
//StringJoin and StringJoinMore

s1 := "hello"
s2 := "world"

//you can use the character which you want join between the strings, using BetweenWith("character")
ret := StringJoin(s1, s2).BetweenWith("") 
fmt.Println(ret)  //"helloworld"

ret := StringJoin(s1, s2).BetweenWith(" ") 
fmt.Println(ret)  //"hello world"

ret := StringJoin(s1, s2).BetweenWith("+") 
fmt.Println(ret)  //"hello+world"

//you also can join more strings like this
StringJoin(s1, s2, s3, s4, ....).BetweenWith("+") 

//you can use array also
str := []string{"hello","world","every"}
ret := StringJoinMore(str).BetweenWith("+")
fmt.Println(ret) //"hello+world"
```
#### SubString
```golang
str := "Monday"
s, _ := SubString(str).Start(0).End(4)
fmt.Println(s) // "Monda"
s, _ : SubString(str).Start(2).End(3)
fmt.Println(s) // "nd"
```
