package main

import "fmt"

func main() {
	example_map := map[int]bool{}
	for i := 0; i < 10; i++ {
		example_map[i] = (i % 2) == 0
	}

	for k, i := range example_map {
		fmt.Printf("key: %d, value: %t \n", k, i)
	}
}
