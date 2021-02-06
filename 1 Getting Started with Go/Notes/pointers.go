package main

import "fmt"

func main(){
  var x = 100
  var y int
  var ip *int

  ip = &x
  y = *ip
  fmt.Println(ip, y)

  ptr := new(int)
  *ptr = 20

  fmt.Println(ptr)

	fmt.Printf("Hello, world!\n")
}
