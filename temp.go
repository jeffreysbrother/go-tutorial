package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"log"
	"math"
	"reflect"
)

func main() {
	fmt.Println(stringutil.Reverse("james"))
	fmt.Println(reflect.TypeOf(66))
	fmt.Println(reflect.TypeOf(1.55))
	fmt.Println(reflect.TypeOf("and then there were none"))
	fmt.Println(reflect.TypeOf(true))

	var a int
	a = 1

	var b, c int
	b = 2
	c = 3

	var d, e = 5, 6

	var f = 7

	g := 8
	h, i := 9, 10
	fmt.Println(a, b, c, d, e, f, g, h, i)

	var wholeNumber int = 1
	var fractionalNumber float64 = 2.5757
	var wholeNumber2 int = int(fractionalNumber)
	var fractionalNumber2 float64 = float64(wholeNumber)
	fmt.Println(wholeNumber2)
	fmt.Println(fractionalNumber2)

	result := add(5, 6)
	fmt.Println(result)

	result2 := difference(12, 5)
	fmt.Println(result2)

	squareRoot, err := squareRoot(9)
	// handle error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(squareRoot)
}

// this function specifies named return type
func add(a float64, b float64) (sum float64) {
	sum = a + b
	// bare return (possible if the return value is named)
	return
}

func difference(a int, b int) int {
	return a - b
}

// multiple return values
func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot do that")
	}
	return math.Sqrt(x), nil
}
