package main

import "fmt"

var Version string = "1.2.3."

func main() {
	// *** Dışarıdaki dosyada erişim için büyük harf ile başlanması lazım. Küçük harf ile başlanırsa erişim yapamaz. ***
	var name string = "golang"

	fmt.Println(name)

}
