---
layout: post
title: "go-defer-example"
date: 2024-03-14
categories: [go 数据结构]
---
这里是一个for range的例子,思考一下会输出什么:
```go
func main() {
  slice := []int{0,1,2,3}
  m := make(map[int]*int)

  for key,val := range slice {
    m[key] = &val
  }

  for k,v := range m {
    fmt.Println(k,"->",*v)
  }
}
```

输出:
```sh
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```
解释：
for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.

正确的写法:
```go
func main() {
  slice := []int{0,1,2,3}
  m := make(map[int]*int)

  for key,val := range slice {
    value := val
    m[key] = &value
  }

  for k,v := range m {
    fmt.Println(k,"===>",*v)
  }
}
```
输出:
```sh
0 ===> 0
1 ===> 1
2 ===> 2
3 ===> 3
```


