package main

import "fmt"

func main() {

	var x [5]int //all the element of the array are initialized to zero when declared
	x[0] = 1 //initializing the first element to 1

	//var y [5]int = [5]{1, 2, 4, 5, 6} // array literal giving an error

	z := [...]int{1, 4, 8, 5}

	z1 := z[1:3]  // slice example
	z2 := z[0:4]

	sli := []int{1,3,4} // slice literal, when defined it creates the underlying array and points to the whole thing

	// sli = make([]int, 10) creating a slice with make method with two arguments type and length/capacity
	// sli = make([]int, 10, 15) will create a slice of size 10 with underlying array of size 15
	// sli = append(sli, 100) will add the element to the slice and can increase the size if necessary

	// MAPS implementation of hash tables in golang

	//var idMap map[string]int // where keys are of the type string and values are of type int
	//idMap := make(map[string]int) // creates a map Using make func
	idMap := map[string]int{
		"joe": 123	}
	idMap["jane"] = 345
	//delete(idMap, "jane") // delete an entry from map using delete with map name and the key
	//id, p := idMap["joe"] // Two value assignment where id will have the value & p will have a boolean ie presence of the key

	fmt.Println(x[1])
	fmt.Println(x[0], "\n")
	//fmt.Println(y[3]) not working
	fmt.Println(z[3], "\n")

	fmt.Println("Using for loop with array")
	for i, v := range z { // range returns both the index and the value of the element
		fmt.Println(i, v)
	}

	fmt.Println(z1[0]) // printing elements from slice
	fmt.Println(z2[0])
	fmt .Println(len(z1), cap(z1)) //array methods length and capacity of the array
	fmt.Println(sli[0])

	fmt.Println("\nMaps")
	fmt.Println(idMap["joe"])
	fmt.Println(len(idMap)) // len function gives the length of the map
	for key, value := range idMap{
		fmt.Println(key, value)
	}
	//fmt.Printf("Hello, world!\n")

	///*
	STRUCTS

	type struct Person {
		name string
		addr string
		phone string
	}

	var p1 Person
	p2 := new(Person)
	p3 := Person(name : "jane", addr : "kaf", phone : "34343221") // Struct literal
	p1.Person = "joe"
	x = p1.addr

}
