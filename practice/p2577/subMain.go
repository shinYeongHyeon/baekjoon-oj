package p2577

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var A, B, C int
	var counts = make([]int, 10)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n%d\n%d\n", &A, &B, &C)

	var multiplyNumber = A * B * C

	for true {
		number := multiplyNumber % 10
		counts[number]++
		multiplyNumber = multiplyNumber / 10

		if multiplyNumber == 0 {
			break
		}
	}

	for i := range counts {
		fmt.Println(counts[i])
	}
}
