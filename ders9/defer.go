package main

import "fmt"

var isConnected bool = false

func main() {

	fmt.Printf("Connections open : %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connections open : %v\n", isConnected)

}

func databaseProcessing() {
	connect()
	fmt.Printf("Defering disconnect!\n")
	defer disconnect()
	fmt.Printf("Connection open :%v\n", isConnected)
}
func connect() {
	isConnected = true
	fmt.Println("Connected to database !")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")
}
