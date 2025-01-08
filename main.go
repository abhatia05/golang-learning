package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var whatToSay string
	whatToSay = "GoodBye, Cool World"
	fmt.Println(whatToSay)
	var i int
	i = 7
	fmt.Println("i has a value", i)
	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
	print(saySomething())
}

func saySomething() string {
	return "My world is you"
}
