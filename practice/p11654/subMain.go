package p11654

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var value string

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &value)

	byteArray := []byte(value)

	fmt.Println(byteArray[0])
}
