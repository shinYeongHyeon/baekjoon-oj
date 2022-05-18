package p1003

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var testCaseCount int
	var zeroAndOneCounts [][]int
	var responses [][]int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &testCaseCount)

	zeroAndOneCounts = append(zeroAndOneCounts, []int{1, 0})
	zeroAndOneCounts = append(zeroAndOneCounts, []int{0, 1})

	for i := 2; i <= 40; i++ {
		zeroAndOneCounts = append(zeroAndOneCounts, []int{
			zeroAndOneCounts[i-1][0] + zeroAndOneCounts[i-2][0],
			zeroAndOneCounts[i-1][1] + zeroAndOneCounts[i-2][1],
		})
	}

	for i := 0; i < testCaseCount; i++ {
		var findIndex int
		fmt.Fscanln(reader, &findIndex)

		responses = append(responses, []int{
			zeroAndOneCounts[findIndex][0],
			zeroAndOneCounts[findIndex][1],
		})
	}

	for _, v := range responses {
		fmt.Println(v[0], v[1])
	}
}
