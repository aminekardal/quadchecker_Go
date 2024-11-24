package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		if i == 0 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					if i == 0 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('C')
					}
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else if i == y-1 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					if i == 0 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('A')
					}
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')

		} else {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}

	}
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadE(x, y)
}
