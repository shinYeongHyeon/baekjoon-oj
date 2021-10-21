package p2438

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
		for j := 1; j <= i; j++ {
			fmt.Fprint(writer,"*")
		}
		fmt.Fprint(writer,"\n")
	}
}
