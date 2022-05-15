package p2480

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var n1, n2, n3 int
	var reward int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n1, &n2, &n3)

	if n1 == n2 && n2 == n3 {
		reward = 10000 + n1*1000
	} else if n1 == n2 {
		reward = 1000 + n1*100
	} else if n2 == n3 {
		reward = 1000 + n2*100
	} else if n1 == n3 {
		reward = 1000 + n3*100
	} else {
		var max = n1

		if n2 > max {
			max = n2
		}

		if n3 > max {
			max = n3
		}

		reward = max * 100
	}

	fmt.Println(reward)
}
