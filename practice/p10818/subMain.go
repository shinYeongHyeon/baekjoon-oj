package p10818

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Exec() {
	var n int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &numbers[i])
	}

	sortAscending(numbers)

	fmt.Printf("%d %d\n", numbers[0], numbers[len(numbers)-1])
}

func sortAscending(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
}
