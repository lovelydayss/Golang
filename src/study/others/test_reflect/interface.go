package main

import (
	"fmt"
)

type m struct {
	ok int
}

type n struct {
	nOk float64
}

func test(u interface{}) {

	// v := reflect.ValueOf(u)

	switch u.(type) {
	case int:
		fmt.Println("is a int")
	case float64:
		fmt.Println("is a float 64")
	case m:
		fmt.Println("is a struct m")
	case n:
		fmt.Println("is a struct n")
	}

}

func main() {

	mm := m{}
	nn := n{}

	test(1)
	test(1.511)
	test(mm)
	test(nn)

}
