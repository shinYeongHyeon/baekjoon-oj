package p8958

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var testNumber int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &testNumber)

	var results = make([]string, testNumber)

	for i := 0; i < testNumber; i++ {
		fmt.Fscanln(reader, &results[i])
		fmt.Println(CalculateScore(results[i]))
	}
}

func CalculateScore(gradingResults string) int {
	var continuousCorrectCount = 0
	var score = 0
	for i := 0; i < len(gradingResults); i++ {
		if gradingResults[i:i+1] == "O" {
			continuousCorrectCount++
			score += continuousCorrectCount
		} else {
			continuousCorrectCount = 0
		}
	}

	return score
}
