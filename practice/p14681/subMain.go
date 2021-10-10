package p14681

import "fmt"

func Exec() {
	var x, y int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Println(getQuadrant(x, y))
}

func getQuadrant(x, y int) int {
	if x > 0 && y > 0 {
		return 1
	} else if x < 0 && y > 0 {
		return 2
	} else if x < 0 && y < 0 {
		return 3
	} else {
		return 4
	}
}
