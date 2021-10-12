package p2739

import "fmt"

func Exec() {
	var num int
	fmt.Scanln(&num)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num * i)
	}
}
