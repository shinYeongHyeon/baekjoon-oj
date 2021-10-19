package p11022

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var length, a, b int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &length)
	defer writer.Flush()

	for i := 1; i <= length; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i, a, b, a + b)
	}
}
