package main

import (
	"
golang.org/x/tour/pic
"
	//"fmt"
	   )

func Pic(dx, dy int) [][]uint8 {
	retSlice := make([][]uint8, dy)
	
	div := 32

	v255 := uint8(255)
	v0 := uint8(0)
	val := v0
	for x := range retSlice {
		data := make([]uint8, dx)
		if x % div == 0 {			
			if val == v255 {
			 		val = v0
				} else {
					val = v255
				}
		}
		for y := range data {
			if y % div == 0 {
				if val == v255 {
			 		val = v0
				} else {
					val = v255
				}
			}

			data[y] = val
		}
		retSlice[x] = data
	}
	return retSlice
}

func main() {
	
pic.Show
(Pic)
}