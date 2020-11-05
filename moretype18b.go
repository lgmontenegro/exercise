package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	retSlice := make([][]uint8, dy)

	for i := range retSlice {
		element := make([]uint8, dx)
		for j := range element {
			if j < 128 && i < 128 {
				element[j] = 0
			} else {
				if j >= 128 && i >= 128 {
					element[j] = 0
				} else {
					element[j] = 255
				}
			}
		}
		retSlice[i] = element
	}

	return retSlice
}

func main() {
	pic.Show(Pic)
}
