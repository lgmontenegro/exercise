package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	retSlice := make([][]uint8, dy)
	element := make([]uint8, dx)
	change := 1
	for i := range retSlice {
		for j := range element {
			switch {
			case change == 1:
				//element[j] = uint8(i^j)
				element[j] = uint8(0)
				change ++
			case change == 2:
				//element[j] = uint8(i*j)
				element[j] = uint8(255)
				change --
			default:
				//element[j] = uint8((i+j)/2)
				element[j] = uint8(change)
			}
		}
		retSlice[i] = element
	}
	return retSlice
}

func main() {
	pic.Show(Pic)
}s