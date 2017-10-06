package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// PowMod function taken from
// http://grokbase.com/p/gg/golang-nuts/16326zgk1v/
func PowMod(a, b, m int64) int64 {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}

// Find powers for given even number
func findMax2Power(x int64) int {
	c := 0

	p := float64(x)

	for math.Mod(p, 2) == 0 {
		c++
		p /= 2
	}

	return c
}

// Implementation of Miller-Rabin primaly test
func MillerRabin(n int64, aces []int64) bool {

	if n == 1 {
		return false
	}

	s := findMax2Power(n - 1)
	d := int64(float64(n-1) / math.Pow(2, float64(s)))

	for _, a := range aces {
		x := PowMod(a, d, n)

		r := 0

		if x == 1 || x == (n-1) {
			continue
		}

		for r = 1; r <= s; r++ {
			x = PowMod(x, 2, n)

			if x == 1 {
				return false // composite
			}

			if x == (n - 1) {
				// continue Aces loop
				r = 0
				break
			}
		}

		if r > 0 {
			return false
		}
	}

	return true
}

func main() {
	{
		aceMap := make([][]int64, 4)
		aceMap[0] = []int64{2}
		aceMap[1] = []int64{2, 3}
		aceMap[2] = []int64{2, 3, 5}
		aceMap[3] = []int64{2, 3, 5, 7}

		var cases [10][2]int64

		reader := bufio.NewReader(os.Stdin)

		firstLine, _, _ := reader.ReadLine()
		numberOfEntries, _ := strconv.Atoi(string(firstLine))

		for n := 0; n < numberOfEntries; n++ {
			{
				line, _, _ := reader.ReadLine()
				x := strings.Split(string(line), " ")

				f, _ := strconv.Atoi(x[0])
				t, _ := strconv.Atoi(x[1])

				cases[n][0] = int64(f)
				cases[n][1] = int64(t)
			}
		}

		for n := 0; n < numberOfEntries; n++ {
			// Sanity checks in case SPOJ assignment assumes
			// corrupted data input
			if cases[n][0] < 1 {
				cases[n][0] = 1
			}

			if cases[n][0] > 999999999 {
				cases[n][0] = 999999999
			}

			if cases[n][1] < 1 {
				cases[n][1] = 1
			}

			if cases[n][1] > 999999999 {
				cases[n][1] = 999999999
			}

			if cases[n][0] != 2 && cases[n][0]%2 == 0 {
				cases[n][0]++
			}

			if cases[n][1] != 2 && cases[n][1]%2 == 0 {
				cases[n][1]--
			}

			if cases[n][0] > cases[n][1] {
				continue
			}

			aces := aceMap[0]

			for i := cases[n][0]; i <= cases[n][1]; i += 2 {
				// Exception handling for 2 being even yet prime number
				if (i == 1 || i == 2) && cases[n][1] >= 2 {
					fmt.Println(2)
					if i == 2 {
						i--
					}
				}

				// Witnesses that give 100% correct ouput as verified by
				// Pomerance, Selfridge, Wagstaff, Jaeschke
				// https://en.wikipedia.org/wiki/Miller–Rabin_primality_test#Deterministic_variants
				if i > 25326000 {
					aces = aceMap[3]
				}

				if i > 1373652 {
					aces = aceMap[2]
				}

				if i > 2046 {
					aces = aceMap[1]
				}

				if MillerRabin(i, aces) {
					fmt.Println(i)
				}
			}

			// Do not print last empty line (for SPOJs sake)
			if n+1 < numberOfEntries {
				fmt.Println()
			}
		}
	}
}
