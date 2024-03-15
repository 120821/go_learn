---
layout: post
title: "go-variable-example"
date: 2024-03-14
categories: [go 语法]
---
这里是一个变量声明的例子,思考一下会输出什么:
```go
var(
  size := 1024
  max_size = size*2
)

func main() {
  fmt.Println(size,max_size)
}
```
输出:
```sh
./variable.go:4:10: syntax error: unexpected :=, expecting =

```
解释：

不能通过编译。这道题的主要知识点是变量声明的简短模式，形如：x := 100。但这种声明方式有限制：

1.必须使用显示初始化；

2.不能提供数据类型，编译器会自动推导；

3.只能在函数内部使用简短模式；

