package p2753

import "fmt"

func Exec() {
	var year int

	fmt.Scanln(&year)
	fmt.Println(CalcLeap(year))
}

func CalcLeap(year int) int {
	if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
		return 1
	} else {
		return 0
	}
}
