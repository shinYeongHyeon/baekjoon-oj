package p10950

import "fmt"

func Exec() {
	var count int
	fmt.Scanln(&count)

	for i := 1; i <= count; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}
