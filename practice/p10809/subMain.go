package p10809

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var inputString string
	var foundIndices []int

	for i := 0; i < 26; i++ {
		foundIndices = append(foundIndices, -1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &inputString)

	for ascii := 97; ascii <= 122; ascii++ {
		var alphabet = string(ascii)
		for i := 0; i < len(inputString); i++ {
			if inputString[i:i+1] == alphabet {
				foundIndices[ascii-97] = i
				break
			}
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%d", foundIndices[i])

		if i == 25 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
