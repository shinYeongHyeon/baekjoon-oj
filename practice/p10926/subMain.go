package p10926

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var id string

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &id)

	fmt.Println(id + "??!")
}
