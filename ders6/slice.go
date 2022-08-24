package main

import "fmt"

func main() {
	/*
		primes := [6]int{1, 2, 3, 4, 5, 6}
		var s []int = primes[0:1]
		fmt.Println(cap(s), len(s))
	*/

	/*
		myArray1 := [...]int{4, 1, 2, 3, 4, 5, 6, 67, 1, 1, 2, 1, 31, 45, 6, 4, 31, 5342, 153, 215, 321, 53, 215}
		mySlice1 := myArray1[:]
		myArray1[0] = 111111
		fmt.Println(mySlice1)
	*/
	/*
		var colors = []string{"red", "green", "bluu"}
		colors = append(colors, "aaa")
		fmt.Println(colors[len(colors)-1])
		fmt.Println(colors[:len(colors)-1])
	*/

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)
	numbers = append(numbers, 3232)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

}
