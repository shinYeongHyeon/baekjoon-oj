package p18108

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var thailanBuddismYear int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &thailanBuddismYear)

	fmt.Println(thailanBuddismYear - 543)
}
