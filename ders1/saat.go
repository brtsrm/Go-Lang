package main

import (
	"fmt"
	"time"
)

func main() {

	// konsol : saat & tarih işlemleri

	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)

	fmt.Printf("Go launch at %s\n", t)

	fmt.Println("------")

	now := time.Now()

	fmt.Printf("The time now is %s\n", now)

	fmt.Println("----")

	fmt.Println("The month is , ", t.Month())
	fmt.Println("The day is , ", t.Day())
	fmt.Println("The weekday is , ", t.Weekday())

	println("----")

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v \n",
		tomorrow.Year(),
		tomorrow.Month(),
		tomorrow.Weekday(),
		tomorrow.Day(),
	)

	println("------")

	longFormat := "Monday, Junuary 2, 2022"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "11/11/06"

	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
