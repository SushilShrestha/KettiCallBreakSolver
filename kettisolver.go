package main

import (
	"fmt"
)

func main() {
	mycard := card{suit: HEARTS, number:10}
	fmt.Println(mycard.number)
}