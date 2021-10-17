package p15552

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {var testLength int

	fmt.Scanln(&testLength)

	var num1, num2 int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < testLength; i++ {
		fmt.Fscanln(reader, &num1, &num2)
		fmt.Fprintln(writer, num1 + num2)
	}

	writer.Flush()
}
