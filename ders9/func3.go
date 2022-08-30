package main

import "fmt"

func main() {

	sayHellos("merhaba", "tellovele", "naber", "dostom")

	sayilar := topla1(1, 2, 3, 4, 5, 6)
	fmt.Println(sayilar)

}

func sayHellos(hellos ...string) {
	for _, hello := range hellos {
		fmt.Println(hello)
	}
}

func topla1(x ...int) int {
	sums := 0
	for _, toplam := range x {
		sums += toplam
	}
	return sums
}
