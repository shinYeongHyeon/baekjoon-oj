package p3052

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	reader := bufio.NewReader(os.Stdin)

	var numbers = make([]int, 10)
	var mods []int
	var divisor = 42

	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])

		var mod = numbers[i] % divisor
		if !isExist(mods, mod) {
			mods = append(mods, mod)
		}
	}

	fmt.Println(len(mods))
}

func isExist(numbers []int, number int) bool {
	for _, n := range numbers {
		if number == n {
			return true
		}
	}

	return false
}
