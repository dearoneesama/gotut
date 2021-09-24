package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dx)
	for x := 0; x < len(p); x++ {
		p[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			p[x][y] = uint8(x^y)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
