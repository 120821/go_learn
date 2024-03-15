package main

func funcMui(x,y int)(sum int,error){
  return x+y,nil
}

func main() {
  funcMui(5, 4)
}
