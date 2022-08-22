package main

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	/*** Örnek 1
	for i, v := range pow {

		fmt.Printf("2**%d = %d\n", i, v)

	} */

	/*** Örnek 2

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	*/

	/*** Örnek Array
	a := [...]string{"a", "b", "c", "d"}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	*/

	/**** Örnek Map
	baskentler := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma", "Japon": "Tokya"}
	for key := range baskentler {
		println("Map item: Captial of ", key, "is", baskentler[key])

	}
	*/

	baskentler := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma", "Japon": "Tokya"}
	for key, val := range baskentler {
		println("Map item: Captial of ", key, "is", val)
	}

}
