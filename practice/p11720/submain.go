package p11720

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Exec() {
	var n, sum int
	var value string

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &value)

	for i := 0; i < n; i++ {
		number, _ := strconv.Atoi(value[i : i+1])
		sum += number
	}

	fmt.Println(sum)
}
