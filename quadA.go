package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)


func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
		z01.PrintRune('\n')
		} else {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
}
}

func main() {
	x,_ := strconv.Atoi(os.Args[1])
	y,_ := strconv.Atoi(os.Args[2])
	QuadA(x,y)
}
