package p1110

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var originNumber int
	var cycle = 0

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &originNumber)
	var newNumber = originNumber

	for true {
		var sum int
		cycle++

		if newNumber < 10 {
			sum = newNumber
		} else {
			sum = newNumber/10 + newNumber%10
		}

		newNumber = (newNumber%10)*10 + sum%10
		if newNumber == originNumber {
			break
		}
	}

	fmt.Println(cycle)
}
