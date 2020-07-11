package main

import "fmt"

func fibonacci(n int) func(int) int {
	cache := make([]int, n+1)
	cache[0] = 0
	cache[1] = 1
	for i:=2; i < n+1; i++ {
			cache[i] = cache[i-1] + cache[i-2]	
	}
	return func(x int) int {	
		return cache[x]
	}
}

func main() {
	var n int = 10
	f := fibonacci(n)
	for i := 0; i < n; i++ {
		fmt.Println(f(i))
	}
}