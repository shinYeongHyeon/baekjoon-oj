package p2742

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

	for i := n; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}

	writer.Flush()
}
