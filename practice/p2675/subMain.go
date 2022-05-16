package p2675

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	reader := bufio.NewReader(os.Stdin)

	var testCount int
	var newStrings []string

	fmt.Fscanln(reader, &testCount)

	for i := 0; i < testCount; i++ {
		var repeatCount int
		var inputString string
		var newString string

		fmt.Fscanln(reader, &repeatCount, &inputString)
		for j := 0; j < len(inputString); j++ {
			for k := 0; k < repeatCount; k++ {
				newString += inputString[j : j+1]
			}
		}

		newStrings = append(newStrings, newString)
	}

	for i := 0; i < len(newStrings); i++ {
		fmt.Println(newStrings[i])
	}
}
