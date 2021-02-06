package main

import "fmt"

func printhello() {
	fmt.Printf("Hello, world!\n")
}

func foo(x int, y int) {
	fmt.Println(x * y)  // parameter of same type can be defined like this (x,y int)
}

func foo1(x int, y int) int {		// int specifies the type of the return value
  return x + y
}

func foo2(x int) (int, int) {
  return x, x + 6735
}

func CAR(y *int)  {		// call by reference by using pointer
  *y = *y + 1
}

func AA(x [3]int) int {		// Array argument - Uses call by value which takes more time if the array is too big
  return x[0]
}

func AA1(x *[3]int) {		// Uses call by reference which takes less time if the array is too big
  (*x)[0] = (*x)[0] + 74
}

func sl(sli []int) {		// instead of call by value and reference we use slice which is like a pointer to kind of copy
  sli[0] = sli[0] + 687
}

func main() {
	//printhello()

	x := 1
	c := [3]int{101, 2, 3}
	d := []int{101, 2, 3}
	foo(4, 84)

	v := foo1(3, 24)
	fmt.Println(v)

	a, b := foo2(7682)
	fmt.Println(a, b)

	CAR(&x)
	fmt.Println(x)

	fmt.Println(AA(c))
	AA1(&c)
	fmt.Println(c)

	sl(d)
	fmt.Println(d)
}
