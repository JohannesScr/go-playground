package example

import (
	"fmt"
	"log"
)

func UserInput() {
	var age int
	fmt.Println("please enter you age:")
	n, err := fmt.Scanln(&age)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Your age is %d\n(%d items successfully scanned)\n", age, n)
}
