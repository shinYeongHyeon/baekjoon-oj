package p3009

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var x1, x2, x3, y1, y2, y3 int
	var nx, ny int
	var reader = bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &x1, &y1)
	fmt.Fscanln(reader, &x2, &y2)
	fmt.Fscanln(reader, &x3, &y3)

	if x1 == x2 {
		nx = x3
	} else if x1 == x3 {
		nx = x2
	} else {
		nx = x1
	}

	if y1 == y2 {
		ny = y3
	} else if y1 == y3 {
		ny = y2
	} else {
		ny = y1
	}

	fmt.Println(nx, ny)
}
