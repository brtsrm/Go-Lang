package main

import "fmt"

func main() {

	// Convert

	/*
		var myString = "1"
		var x = 10
		var f float32 = 2.0

		fmt.Println(myString, x, f)

		num, _ := strconv.Atoi(myString)
		fmt.Println("Sonuc : " + strconv.Itoa(num))
		result := num + 2
		fmt.Println(result)
	*/
	// CASTING

	var i int = 55
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

}
