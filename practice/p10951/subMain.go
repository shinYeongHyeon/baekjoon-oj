package p10951

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var a, b int
	var requiredInputCount = 2

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		val, _ := fmt.Fscanln(reader, &a, &b)
		if val != requiredInputCount {
			break
		}

		fmt.Fprintln(writer, a+b)
	}
}
