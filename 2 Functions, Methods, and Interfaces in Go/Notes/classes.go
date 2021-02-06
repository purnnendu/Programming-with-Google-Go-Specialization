package main

import "fmt"
import "math"

type MyInt int

type Point struct {
	x float64
	y float64
}

func (mi MyInt) double() int {
	return int(mi * 2)
}

func (p Point) DistToOrigin () {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {

	//fmt.Printf("Hello, world!\n")

	v := MyInt(3)
	fmt.Println(v.double())  // only a copy of V is passed to MyInt.double(call by value) and computed is printed
	fmt.Println(v)

	p1 := Point(3, 4)
	fmt.Println(p1.DistToOrigin())

}

// might need to review
