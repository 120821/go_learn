---
layout: post
title: "go-return-example"
date: 2024-03-14
categories: [go 语法]
---
这里是一个return的例子,思考一下会输出什么:
```go
func funcMui(x,y int)(sum int,error){
  return x+y,nil
}
```

输出:

会panic

```sh
syntax error: mixed named and unnamed parameters
```
解释：

在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。
如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
这里的第一个返回值有命名 sum，第二个没有命名，所以错误。


