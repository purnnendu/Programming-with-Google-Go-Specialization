package main

import "fmt"

func main() {
	fmt.Printf("Hello, world!\n")
	f()
	g()
	f1()
	h() 
}

var x = 2

func f(){
	fmt.Printf("%d",x)
}

func g(){
	fmt.Printf("%d",x)
}

func f1(){
	var y = 2
	fmt.Printf("%d",y)
}

func h(){
	fmt.Printf("%d",y)
}
