package main

import (
	"fmt"

	"../utils"
)

func main() {
	n1, l1 := utils.FullName("Merhaba", "Televole")
	fmt.Printf("fullnam: %v number char: %v", n1, l1)

	n2, l2 := utils.FullNameNakedReturn("Merhaba", "Televole")
	fmt.Printf("fullnam: %v number char: %v", n2, l2)
}
