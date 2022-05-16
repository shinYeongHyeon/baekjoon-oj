package p1157

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exec() {
	reader := bufio.NewReader(os.Stdin)

	var word string
	var alphabetCounts [26]int
	var max, maxIndex int
	var isDuplicated = false

	fmt.Fscanln(reader, &word)
	word = strings.ToLower(word)
	asciis := []rune(word) // Output: [97 98 99]

	for _, ascii := range asciis {
		alphabetCounts[ascii-97]++
	}

	for i, value := range alphabetCounts {
		if value > max {
			isDuplicated = false
			max = value
			maxIndex = i
		} else if value == max {
			isDuplicated = true
		}
	}

	if isDuplicated {
		fmt.Println("?")
	} else {
		fmt.Println(string(maxIndex + 65))
	}
}
