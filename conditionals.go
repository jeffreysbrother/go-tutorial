package main

import "fmt"

func main() {
	if 1 < 2 {
		fmt.Println("one is less than 2")
	} else if true || false {
		fmt.Println("true || false")
	}
}
