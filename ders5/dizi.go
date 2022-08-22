package main

import "fmt"

func main() {

	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54

	fmt.Println(myArray1)

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "Purple"
	fmt.Println(colors[1])

	var numbers = [5]int{21, 211, 32, 321, 22}
	fmt.Println(numbers)
	fmt.Println("Number of numbers", len(numbers))

	myArray2 := [...]int{4, 3, 2, 1, 3}
	fmt.Println(len(myArray2), myArray2)

	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Volvo"
	cars[2] = "Lamborgin"

	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	// While
	fmt.Println("While")

	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}

	fmt.Println("For")
	// Len
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	// Range
	fmt.Println("Range")
	for i := range cars {
		fmt.Println(cars[i])
	}

}
