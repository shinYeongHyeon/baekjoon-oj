package p4673

import "fmt"

func Exec() {
	var createdNumbers []bool

	for i := 0; i < 10000; i++ {
		createdNumbers = append(createdNumbers, false)
	}

	var lastCreatedNumber = 0
	for i := 0; i < 10000; i++ {
		lastCreatedNumber = d(i)
		if lastCreatedNumber >= 10000 {
			continue
		}

		createdNumbers[lastCreatedNumber] = true
	}

	for i := 0; i < len(createdNumbers); i++ {
		if !createdNumbers[i] {
			fmt.Println(i)
		}
	}
}

func d(n int) int {
	if n < 10 {
		return n + n
	}

	if n < 100 {
		return n + n/10 + n%10
	}

	if n < 1000 {
		return n + n/100 + (n/10)%10 + n%10
	}

	if n < 10000 {
		return n + n/1000 + (n/100)%10 + (n/10)%10 + n%10
	}

	return 10001
}
