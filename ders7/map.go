package main

import "fmt"

func main() {
	/*
		myMap := make(map[int]string)
		fmt.Println(myMap)
		myMap[43] = "Foo"
		myMap[12] = "Bar"
		fmt.Println(myMap)
	*/

	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["DYB"] = "Diyarbakır"
	fmt.Println(states)

	diyarbakir := states["DYB"]

	fmt.Println("Seçilen şehir", diyarbakir)
	delete(states, "DYB")
	fmt.Println(states)
	states["ANT"] = "Antalya"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("Key : %v Value : %v \n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("\n", keys)

	// Key listesinde indeks değerlerine göre states nesnesinde bulunan şehileri yazdır

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
