---
layout: post
title: "go-new-and-make"
date: 2024-03-14
categories: [go 数据结构]
---

new 的作用是初始化一个内置类型的指针 new 函数是内建函数，函数定义：
```go
func new(Type) *Type
```

1.使用 new 函数来分配空间

2.传递给 new 函数的是一个类型，而不是一个值

3.返回值是指向这个新分配的地址的指针

make 的作用是为 slice, map or chan 的初始化 然后返回引用
make 函数是内建函数，函数定义：
```go
func make(Type, size IntegerType) Type
```
make(T, args)函数的目的和 new(T)不同 仅仅用于创建 slice, map,
channel 而且返回类型是实例

