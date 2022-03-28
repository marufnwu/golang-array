package main

import (
	"fmt"
	"strings"
)

func main() {
	var number [4]int

	fmt.Printf("%v\n", number) //print only elements[0 0 0 0]

	fmt.Printf("%#v\n", number) //print type with elements [4]int{0, 0, 0, 0}
	
	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	a2 := [4]string{"x", "y"}
	fmt.Printf("%v\n", a2)

	//ellipsis operator
	a5 := [...]int{1,3} //dynamic size
	fmt.Printf("Size of array %d", len(a5))

	a6 := [...]int{
		1,
		2,
		3,
		4, //this comma is mandatory for multiline initilization
	}

	_ = a6
	//iterate array
	for i, v := range a5{
		fmt.Printf("\n%v position is %v", i, v)
	}

	fmt.Println("\n",strings.Repeat("#", 20))

	for i := 0; i < len(a5); i++ {
		fmt.Printf("\n%v position is %v", i, a5[i])
	}

	fmt.Println("\n",strings.Repeat("#", 20))
	//multi dimentional array

	balance := [3][3]int{
		{2,3,4},
		{5,6,7},
		
	}

	fmt.Println(balance)

	//array item declear with index or position

	_ = [3]int{
		1: 0,
		2: 4,
		0: 5,
	}

	array := [4]int{
		1,
		2,
		3,
	}

	fmt.Println("Array: ",array)

	slice := []int{
		1,
		2,
		3,
	}

	fmt.Println("Slice: ",slice)

}
