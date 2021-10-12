package p8393

import "fmt"

func Exec() {
	var num int
	var sum int = 0
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
