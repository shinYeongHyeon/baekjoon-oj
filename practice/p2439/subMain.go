package p2439

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var number int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &number)
	defer writer.Flush()

	for i := 1; i <= number; i++ {
		for j := number - i - 1; j >= 0; j-- {
			fmt.Fprint(writer, " ")
		}
		for k := 1; k <= i; k++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
}
