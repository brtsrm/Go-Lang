package main

import "fmt"

func main() {

	addFunc := func(terms ...int) (numTers int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTers = len(terms)
		return numTers, sum
	}

	fmt.Println(addFunc(1, 2, 3, 4, 5, 6, 7, 8, 9))

}
