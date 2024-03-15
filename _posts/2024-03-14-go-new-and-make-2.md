---
layout: post
title: "go-new-and-make-2"
date: 2024-03-14
categories: [go 基础知识]
---
1. new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

2. new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。
换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

3. make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。
make() 只适用于 slice、map 和 channel.

这里是一个new之后append的例子,思考一下会输出什么:
```go
func main() {
  list := new([]int)
  list = append(list, 1)
  fmt.Println(list)
}
```

输出:
```sh
first argument to append must be a slice; have list (variable of type *[]int)
```
解释：

不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。
可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

