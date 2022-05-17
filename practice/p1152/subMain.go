package p1152

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exec() {
	var count int
	var inputString string

	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')

	words := strings.Split(inputString, " ")
	for i := range words {
		if words[i] == "\n" || words[i] == "" {
			continue
		}

		count++
	}

	fmt.Println(count)
}
