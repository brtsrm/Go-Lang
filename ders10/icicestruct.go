package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Kullanıcı veri oluşturma alanı

	fmt.Println("Kullanıcı oluşturma v1")

	users1 := &User{
		ID:        1,
		firstName: "Berat",
		lastName:  "Sarmış",
		age:       25,
		pay: &Payment{
			Salary: 3333,
			Bonus:  600,
		},
	}

	fmt.Println("Maaş : " + strconv.FormatFloat(users1.GetPayment(), 'g', -1, 64))

}

// kullanıcı yapısı

type Payment struct {
	Salary float64
	Bonus  float64
}

func newPaymnet() *Payment {
	p := new(Payment)
	return p
}

type User struct {
	ID        int
	firstName string
	lastName  string
	userName  string
	age       int
	pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.pay = newPaymnet()
	return u
}

func (u User) GetFullName() string {
	return u.firstName + " " + u.lastName
}
func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetPayment() float64 {
	pay := u.pay.Salary + u.pay.Bonus
	return pay
}
