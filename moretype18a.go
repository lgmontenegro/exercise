package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	retSlice := make([][]uint8, dy)
	azulao := make([]uint8, dx)
	branquinha := make([]uint8, dx)

	for j := range azulao {
		if j < 128 {
			azulao[j] = 0
			branquinha[j] = 255
		} else {
			azulao[j] = 255
			branquinha[j] = 0
		}

	}

	for i := range retSlice {
		if i < 128 {
			retSlice[i] = azulao
		} else {
			retSlice[i] = branquinha
		}
	}

	return retSlice
}

func main() {
	pic.Show(Pic)
}
