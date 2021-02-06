package main

import "fmt"

func foo() *int {
	x := 100
	return &x
}

func main() {
	var y *int
	y = foo()
	fmt.Println(*y) // why deallocation is difficult

	var applenum int
	fmt.Println("Enter the number of the apple?")
	num, err := fmt.Scan(&applenum)
	fmt.Println(num, err, "The number of apple is", applenum)
}

/* this
is
a
block
of
multi line
comment*/

/*

if condition {

}else {

}

for index := 0;  < count; ++ {

}

i := 0
for i < count; {
	fmt.Println("high")
	i++
}

for {
	fmt.Println("infinite loop")
}

switch expression {
case 1:fmt.Println(1)
case 2:fmt.Println(2)
default:fmt.Println("default")
}
*/
