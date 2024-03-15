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
