package main

import "fmt"

func main() {

	// fmt.Println(add(12, 21))
	sum, numberLen := add(3, 4, 5)

	fmt.Println(numberLen, sum)

}

/*
func sayHello(message *string) {
	println(*message)
	*message = "Merhaba"
}
*/
/*
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
*/

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result, len(terms)
}
