package p1546

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	// N + A, B, C
	// (A / A * 100 + B / A * 100 + C / A * 100) / N
	// (A + B + C) / A / N * 100

	var number int
	var max = 0
	var sum = 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &number)

	var scores = make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Fscanf(reader, "%d", &scores[i])
		if max < scores[i] {
			max = scores[i]
		}
		sum += scores[i]
	}

	fmt.Println(float64(sum) / float64(max) / float64(number) * 100)
}
