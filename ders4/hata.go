package main

import (
	"log"
	"os"
)

func main() {

	/***** Örnek 1
	file, err := os.Open("dosyam.txt")

	if err != nil {
		os.Create("pok.txt")
	}
	*/

	/***** Örnek 2
		myError := errors.New("Bu bir hata !")
		fmt.Println(myError.Error())
	***/

	/***** Örnek3
	i := -2

	if i < 0 {
		return 0, fmt.Errorf("Matematik : çok kötü bir hata %g", i)
	}
	*/

	_, err := os.Open("abc.rar")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		// fmt.Println("Error : ", err.Error())
		// os.Exit(1)
		log.Fatalln("Error : ", err)
	}
}
