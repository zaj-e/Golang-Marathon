//(nextPage)
// You can skip the index or value by assigning to _. 
// for i, _ := range pow
// for _, value := range pow
// If you only want the index, you can omit the second variable.
// for i := range pow

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
