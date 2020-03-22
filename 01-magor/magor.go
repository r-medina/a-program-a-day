// everything base 10

// TODO: probably remove decimals

// TODO: problem with '...' in zshell

package main

import (
	"fmt"
	"os"
	"strconv"
)

// digitToSound maps a number to a broad phonetic transcription.
var digitToSound = map[int][]string{
	0: []string{"s", "z"},
	1: []string{"t", "d"},
	2: []string{"n"},
	3: []string{"m"},
	4: []string{"r" /* but also l as in colonel*/},
	5: []string{"l"},
	6: []string{"tʃ", "dʒ", "ʃ", "ʒ"},
	7: []string{"k", "g"},
	8: []string{"f", "v"},
	9: []string{"p", "b"},
}

// StringToWords takes a string containing numbers and returns a list of words
// that can be used to memorize the numbers
func StringToWords(str string) ([]string, error) {
	return nil, nil
}

func stringToSounds(s string) ([][]string, error) {
	digits, err := stringToDigits(s)
	if err != nil {
		return nil, fmt.Errorf("error getting digits from %q: %v", s, err)
	}

	return digitsToSounds(digits)
}

func stringToDigits(s string) ([]int, error) {
	digits := []int{}
	for i, c := range s {
		if c == '.' || c == ',' || c == '-' {
			continue
		}
		d, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, fmt.Errorf("error converting %q digit %d: %v", c, i, err)
		}
		digits = append(digits, d)
		i /= 10
	}

	return digits, nil

}

func digitsToSounds(digits []int) ([][]string, error) {
	s := make([][]string, len(digits))

	for i, d := range digits {
		if d > 9 || d < 0 {
			return nil, fmt.Errorf("digit %d invalid value: %d", i, d)
		}
		s[i] = digitToSound[d]
	}

	return s, nil
}

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func printSounds(sounds [][]string) {
	for _, s := range sounds {
		fmt.Print("[")
		isFirst := true
		for _, idk := range s {
			if !isFirst {
				fmt.Print(" ")
			}
			fmt.Printf("/%s/", idk)
			isFirst = false
		}
		fmt.Print("] ")
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: magor NUM")
		os.Exit(0)
	}
	sounds, err := stringToSounds(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printSounds(sounds)
}
