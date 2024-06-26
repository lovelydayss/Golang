package main

import "fmt"

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals

			if !ok {
				break
			}
			squares <- x * x
		}

		close(squares)
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
