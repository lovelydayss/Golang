package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func producer(wg *sync.WaitGroup, data chan<- int) {

	defer wg.Done()

	nums := []int{1, 2, 3, 4, 5, 6}

	for _, num := range nums {
		data <- num
	}

	close(data)

}

func consumer(wg *sync.WaitGroup, data <-chan int) {

	defer wg.Done()

	for num := range data {

		fmt.Printf("value is %d \n", num)

	}

}

func main() {

	// golang 指针默认 8 byte
	var x *string
	fmt.Println(unsafe.Sizeof(x))

	data := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	// producer
	go func() {

		defer wg.Done()

		nums := []int{1, 2, 3, 4, 5, 6}

		for _, num := range nums {
			data <- num
		}

		close(data)

	}()

	// consumer
	go func() {

		defer wg.Done()

		for num := range data {

			fmt.Printf("value is %d \n", num)

		}

	}()

	wg.Wait()

}
