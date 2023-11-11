package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	A := 0.0
	B := 0.0
	z := make([]float64, 1760)
	b := make([]byte, 1760)

	fmt.Print("\033c")
	fmt.Print("\x1b[2J")

	for {
		for i := range b {
			b[i] = ' '
		}
		for i := range z {
			z[i] = 0
		}

		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				h := math.Cos(j) + 2
				D := 1 / (math.Sin(i)*h*math.Sin(A) + math.Sin(j)*math.Cos(A) + 5)

				t := math.Sin(i)*h*math.Cos(A) - math.Sin(j)*math.Sin(A)

				x := int(40 + 30*D*(math.Cos(i)*h*math.Cos(B)-t*math.Sin(B)))
				y := int(12 + 15*D*(math.Cos(i)*h*math.Sin(B)+t*math.Cos(B)))
				o := x + 80*y
				N := 8 * ((math.Sin(j)*math.Sin(A)-math.Sin(i)*math.Cos(j)*math.Cos(A))*math.Cos(B) -
					math.Sin(i)*math.Cos(j)*math.Sin(A) - math.Sin(j)*math.Cos(A) -
					math.Cos(i)*math.Cos(j)*math.Sin(B))

				if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
					z[o] = D
					if N > 0 {
						b[o] = ".-:!*oe&#%&@"[int(math.Floor(N))]
					} else {
						b[o] = '.'
					}
				}
			}
		}

		fmt.Print("\x1b[H")
		for y := 0; y < 22; y++ {
			for x := 0; x < 80; x++ {
				fmt.Printf("%c", b[x+80*y])
			}
			fmt.Println()
		}

		A += 0.04
		B += 0.02

		time.Sleep(10 * time.Millisecond)
	}
}
