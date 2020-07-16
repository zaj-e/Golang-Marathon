package main

import (
	"fmt"
	"github.com/zaj-e/lib/string"
	"github.com/zaj-e/lib/sort"
)

func main(){
	fmt.Println(string.Reverse("Hello world"))
	m := []int{8,2,5,4}
	fmt.Println(m)
	sort.Sorter(m)
	fmt.Println(m)
}