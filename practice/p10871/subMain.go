package p10871

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var n, x int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &x)
	defer writer.Flush()

	// sequence : 수열
	var sequence = make([]int, n)
	for i := range sequence {
		fmt.Fscanf(reader, "%d", &sequence[i])
		if sequence[i] < x {
			fmt.Fprintf(writer, "%d ", sequence[i])
		}
	}
	fmt.Fprint(writer, "\n")
}
