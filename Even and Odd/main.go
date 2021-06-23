package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, number := range ints {

		if ints[i]%2 != 0 {
			fmt.Println(number, "is odd")
		} else {
			fmt.Println(number, "is even")
		}
	}

}
