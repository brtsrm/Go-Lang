package main

import "fmt"

func main() {

	var puan float64
	fmt.Println("Sınav notunuzu giriniz :")
	fmt.Scanf("%v", &puan)
	switch {
	case puan <= 44:
		fmt.Println("1 aldınız")
	case puan <= 54:
		fmt.Println("2 aldınız")
	case puan <= 69:
		fmt.Println("3 aldınız")
	case puan <= 84:
		fmt.Println("4 aldınız")
	case puan <= 100:
		fmt.Println("5 aldınız")
	}

}
