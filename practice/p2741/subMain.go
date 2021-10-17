package p2741

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var n int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, i+1)
	}

	writer.Flush()
}
