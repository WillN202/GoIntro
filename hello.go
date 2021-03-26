// The first line in a go file is the package name, all files belong to a package. If we want to run the package it needs to be called main
package main

//This is followed by imports for other packages e.g. fmt for IO. There are no commas in this list
import (
	"errors"
	"fmt"
	"math"
)

//Strucs are useful for holding collections of fields for a more logical type. Type <name> struct
//Jump to near the bottom of main() for how they are used

type person struct {
	//Each field names a name and type or even other structs
	name string
	age  int
}

// this is the entry point
func main() {
	fmt.Println("hello world")

	//variables have the format var <name> <type>. They will have many zero values
	var x int
	x = 5
	var y int = 7
	sum := x + y
	fmt.Println(x)

	if x > y {
		fmt.Println("X more than 6")
	} else if sum < 5 {
		fmt.Println("Sum less than 5")
	} else {
		fmt.Println("no cond satisifed")
	}

	// Arrays must have the same type and are fixed and are 0 indexed

	var a [5]int
	fmt.Println(a)
	a[2] = 5
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	//Slices are an abstraction over an array with no fixed length, this is done by removing the element count

	c := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(c)
	c = append(c, 13)
	//maps are like dicts with the types of keys then values

	vert := make(map[string]int)
	vert["triangle"] = 3
	vert["square"] = 4

	fmt.Println(vert)
	fmt.Println(vert["triangle"])

	delete(vert, "triangle")
	fmt.Println(vert)

	// go only supports for loops, variable, stopping condition and looping action buy by deleting the variable and increment this will become a while loop

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----")
	for j := 0; j < 3; {
		fmt.Println(j)
		j++
	}

	// We can use range to create for each loops over arrays/slices and maps
	arr := []string{"a", "b", "c", "d"}

	for index, value := range arr {
		fmt.Println("index: ", index, "value", value)
	}

	for key, value := range vert {
		fmt.Println("key: ", key, "value", value)
	}

	fmt.Println(sums(2, 3))

	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//Now we create a struct of the type
	p := person{name: "Name", age: 19}
	fmt.Println(p)
	fmt.Println(p.age)

	//Go also supports pointers &<variablename> is the pointer to that variable
	m := 7
	fmt.Println(m, &m)
	sum2(&m)
	fmt.Println(m, &m)

}

// Functions are declared, with the name, (arguments) return type. Funcs can have multiple return types. They will always pass by value unless you use pointers
func sums(x int, y int) int {
	return x + y
}

//Nill == None, this allows us to restrict functions to certain values. This is because Go has no concept of exceptions
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Cannot sqrt negative")
	}

	return math.Sqrt(x), nil
}

//To accept a pointer we put * before the type like in C. * is also used to derefeence pointer
func sum2(x *int) {
	*x++
}

// to run this type go run <filename>, to make it into an executeable type go build. Remember your project must be in a directory for this to work
//go install does something similar but will put the file into a bin folder, if there were externald depencies they could be cached in another folder
