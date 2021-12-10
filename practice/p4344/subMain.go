package p4344

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var testNumber int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &testNumber)
	defer writer.Flush()

	for i := 0; i < testNumber; i++ {
		var studentNumber int
		fmt.Fscanf(reader, "%d ", &studentNumber)

		var scores = make([]float64, studentNumber)
		var sum, avg float64
		for j := 0; j < studentNumber; j++ {
			fmt.Fscanf(reader, "%f ", &scores[j])
			sum += scores[j]
		}
		avg = sum / float64(studentNumber)

		var aboveAverageCount float64
		for _, val := range scores {
			if val > avg {
				aboveAverageCount += 1
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", aboveAverageCount/float64(studentNumber)*100, "%")
	}
}
