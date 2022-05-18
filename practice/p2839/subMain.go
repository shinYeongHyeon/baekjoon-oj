package p2839

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var kg int
	var bag int
	var reader = bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &kg)

	for ; kg >= 0; kg -= 3 {
		if kg%5 == 0 {
			bag += kg / 5
			fmt.Println(bag)
			return
		}

		bag++
	}

	fmt.Println(-1)
}
