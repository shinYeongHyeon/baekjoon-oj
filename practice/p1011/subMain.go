package p1011

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var testCaseCount int
	var responses []uint64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &testCaseCount)

	for i := 0; i < testCaseCount; i++ {
		var x, y, k uint64
		fmt.Fscanln(reader, &x, &y)
		var distance = y - x

		for k = 1; true; k++ {
			if distance <= k*k {
				responses = append(responses, k*2-1)
				break
			} else if distance <= k*k+k {
				responses = append(responses, k*2)
				break
			}
		}
	}

	for i := 0; i < len(responses); i++ {
		fmt.Println(responses[i])
	}
}
