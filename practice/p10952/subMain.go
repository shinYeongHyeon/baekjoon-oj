package p10952

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var a, b int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		fmt.Fscanln(reader, &a, &b)

		if a == 0 && b == 0 {
			break
		}

		fmt.Fprintln(writer, a+b)
	}
}
