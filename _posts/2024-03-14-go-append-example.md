---
layout: post
title: "go-append-example"
date: 2024-03-14
categories: [go 语法]
---
这里是一个new之后append的例子,思考一下会输出什么:
```go
func main() {
  s1 := []int{1, 2, 3}
  s2 := []int{4, 5}
  s1 = append(s1, s2)
  fmt.Println(s1)
}
```

输出:
```sh
cannot use s2 (variable of type []int) as type int in argument to append
```
解释：

不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。
或者直接跟上元素，形如：append(s1,1,2,3)。

