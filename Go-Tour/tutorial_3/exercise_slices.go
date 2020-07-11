package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var matrix [][]uint8
	for i := 0; i < dy; i++ {
		var row []uint8
		for j := 0; j < dy; j++ {
			row = append(row, uint8(transform(i,j)))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func transform (x int, y int) int {
	//return (x+y)/2
	return (x*y)
}


func main() {
	pic.Show(Pic)
}
