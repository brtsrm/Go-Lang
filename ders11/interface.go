package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}
func (_ Ferrari) Stop() bool {
	return false
}
func (x *Ferrari) Information() string {
	ret := "Marka : " + x.Brand + " \n Model : " + x.Model + " \n  Color" + x.Color + " \n  Speed" + strconv.Itoa(x.Speed) + " \n  Price" + strconv.FormatFloat(x.Price, 'g', -1, 64) + "€"
	add := "Yes"
	if x.Special {
		ret += "\n " + "\t" + "Sepical" + add
	}
	return ret
}

type Lamborghini struct {
	Car
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}
func (_ Lamborghini) Stop() bool {
	return true
}
func (x *Lamborghini) Information() string {
	ret := "Marka : " + x.Brand + " \n Model : " + x.Model + " \n  Color" + x.Color + " \n  Speed" + strconv.Itoa(x.Speed) + " \n  Price" + strconv.FormatFloat(x.Price, 'g', -1, 64) + "€"
	add := "Yes"
	if x.Special {
		ret += "\n " + "\t" + "Sepical" + add
	}
	return ret
}

type Mercedes struct {
	Car
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}
func (_ Mercedes) Stop() bool {
	return true
}
func (x *Mercedes) Information() string {
	ret := "Marka : " + x.Brand + " \n Model : " + x.Model + " \n  Color" + x.Color + " \n  Speed" + strconv.Itoa(x.Speed) + " \n  Price" + strconv.FormatFloat(x.Price, 'g', -1, 64) + "€"
	add := "Yes"
	if x.Special {
		ret += "\n " + "\t" + "Special" + add
	}
	return ret
}
func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Araç bilgi : " + c.Information())
	fmt.Println("\n")
	msg := ""
	isRun := c.Run()
	if isRun {
		msg = "Çalışıyor"
	} else {
		msg = "Çalışmıyor"
	}
	fmt.Println("Araç bilgisi " + msg + ".")

	isStop := c.Stop()
	if isStop {
		fmt.Println("durdu")
	} else {
		fmt.Println("duramadı fren tutmuyor")
	}

}
func main() {
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "f20"
	ferr.Color = "Red"
	ferr.Speed = 300
	ferr.Price = 1.4
	ferr.Special = true

	fmt.Println(ferr.Information())
	CarExecute(ferr)

}
