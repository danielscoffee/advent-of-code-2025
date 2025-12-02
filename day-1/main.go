package main

import (
	"fmt"
	"os"
)

const (
	L = 0
	R = 1
)

type Keyword struct {
	turns     int
	direction int
}

func (k *Keyword) Clocking(key []Keyword) int {
	position := 50
	zeroCount := 0

	for i := range key {
		turns := key[i].turns
		for j := 0; j < turns; j++ {
			if key[i].direction == L {
				position--
				if position < 0 {
					position = 99
				}
			}
			if key[i].direction == R {
				position++
				if position > 99 {
					position = 0
				}
			}
		}
		if position == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func (k *Keyword) ZeroCounters(key []Keyword) int {
	position := 50
	zeroCount := 0

	for i := range key {
		turns := key[i].turns
		for j := 0; j < turns; j++ {
			if key[i].direction == L {
				position--
				if position < 0 {
					position = 99
				}
			}
			if key[i].direction == R {
				position++
				if position > 99 {
					position = 0
				}
			}
			if position == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}

func stringToKeyword(s string) []Keyword {
	ka := []Keyword{}

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' || s[i] == 'R' {
			direction := L
			if s[i] == 'R' {
				direction = R
			}
			i++
			numStart := i
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			if numStart < i {
				num := 0
				for j := numStart; j < i; j++ {
					num = num*10 + int(s[j]-'0')
				}
				ka = append(ka, Keyword{turns: num, direction: direction})
			}

			i--
		}
	}
	return ka
}

func main() {
	d1, err := os.ReadFile("input_part1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	d2, err := os.ReadFile("input_part2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	k1 := stringToKeyword(string(d1))
	ka := Keyword{}
	k2 := stringToKeyword(string(d2))

	password1 := ka.Clocking(k1)
	fmt.Printf("Password: %d\n", password1)
	password2 := ka.ZeroCounters(k2)
	fmt.Printf("Password: %d\n", password2)
}
