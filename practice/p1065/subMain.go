package p1065

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	// NOTE: 어떤 양의 정수 X의 각 자리가 등차수열을 이룬다면, 그 수를 한수라고 한다.
	var n int
	var count int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	for i := 1; i <= n; i++ {
		if isHansoo(i) {
			count++
		}
	}

	fmt.Println(count)
}

func isHansoo(value int) bool {
	if value < 100 {
		return true
	}

	if value == 1000 {
		return false
	}

	var firstDiff = value/100 - value/10%10
	var secondDiff = value/10%10 - value%10

	if firstDiff == secondDiff {
		return true
	}

	return false
}
