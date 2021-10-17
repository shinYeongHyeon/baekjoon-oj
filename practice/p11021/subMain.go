package p11021

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var testLength int

	fmt.Scanln(&testLength)

	var num1, num2 int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 1; i <= testLength; i++ {
		fmt.Fscanln(reader, &num1, &num2)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, num1 + num2)
	}

	writer.Flush()
}
