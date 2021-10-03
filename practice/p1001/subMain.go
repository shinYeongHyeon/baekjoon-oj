package p1001

import "fmt"

func Exec() {
	var num1 int
	var num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Printf("%d", num1 - num2)
}
