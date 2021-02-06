package main

import "fmt"
import "math"

var funcVar func(int) int  // declaring function variable called funcVar

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int {
	return x + 1
}

func decFn(x int) int {return x - 1}

func MakeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
	// Example for returning function ie a func to make different func based on differnet parameters
	fn := func (x, y float64) float64 {
		return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
	}
	return fn
}

func getMax(vals ...int) int {
	// Example for passing variable number of arguments ie you pass it as a slice
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {

	//fmt.Printf("Hello, world!\n")

	defer fmt.Println("Bye!")
	// defer is used for executing something at the end like a cleanup. You can even use func with defer

	funcVar = incFn
	/* Assigning funcVar to the func incFn ie funcVar(23) & incFn(23) would be same & incFn func can be accessed with funcVar
	variable, its called treating the func as a first class object */
 	fmt.Println(funcVar(21))

	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	v := applyIt(func (x int) int {return x + 1}, 43)
	// Anonymous func ie a func with no name. its also calles lambda func similar to py
	fmt.Println(v)

	dist1 := MakeDistOrigin(0, 0)
	dist2 := MakeDistOrigin(2, 2)
	fmt.Println(dist1(2, 2))
	fmt.Println(dist2(2, 2))

	fmt.Println(getMax(1, 4, 3 , 24378))
	vslice := []int{1, 43, 32403, 6, 46}
	fmt.Println(getMax(vslice...))

	i := 0
	defer fmt.Println(i + 1)
	fmt.Println(i + 1)

	fmt.Printf("End\n")

}
