package main

/*
package main is a convention to tell the go tool to produce an executable
command instead of a package object that would be imported by other code
*/

import (
	"fmt"
	"unique_namespace/string"
)

func main(){
	fmt.Println(string.Reverse("Hello"))
}