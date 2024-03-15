---
layout: post
title: "go-defer-example-2"
date: 2024-03-14
categories: [go 语法]
---
这里是一个defer使用的例子,思考一下会输出什么:
```go
package main
func test1()bool{
  a := false
  defer func(){
    a = true
  }
  return a
}
func test2()(a bool){
  a = false
  defer func(){
    a = true
  }
  defer func() {
    a = true
  }()
  return a
}
```
输出:
```sh
./defer_example.go:4:11: expression in defer must be function call
./defer_example.go:11:11: expression in defer must be function call
```
解释：
defer 语句后面的 ( ) 是用于表示推迟执行的是一个函数调用，而不是一个表达式。

如果省略 ( )，那么编译器会认为后面的内容是一个表达式，而不是一个函数调用，因此会报语法错误。
正确的写法是:
```go

package main
func test1()bool{
  a := false
  //  defer 语句中使用了匿名函数，并且在匿名函数中修改了变量 a 的值。
  // 然而，在 Go 语言中，defer 语句要求后面跟随的是一个函数调用，而不是一个表达式。
  //defer func(){
  //  a = true
  //}
  defer func() {
    a = true
  }()
  return a
}
func test2()(a bool){
  a = false
  //  defer 语句中使用了匿名函数，并且在匿名函数中修改了变量 a 的值。
  // 然而，在 Go 语言中，defer 语句要求后面跟随的是一个函数调用，而不是一个表达式。
  //defer func(){
  //  a = true
  //}
  defer func() {
    a = true
  }()
  return a
}
func main() {
  print(test1()) // false
  print(test2()) // true
}
```


