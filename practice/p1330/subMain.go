package p1330

import "fmt"

func Exec() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a > b {
		fmt.Println(">")
	} else if a < b {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}
