package p2884

import "fmt"

func Exec() {
	var hour, minute int
	fmt.Scanln(&hour, &minute)
	if minute - 45 >= 0 {
		fmt.Printf("%d %d", hour, minute - 45)
	} else if hour == 0 {
		fmt.Printf("%d %d", 23, minute + 15)
	} else {
		fmt.Printf ("%d %d", hour - 1, minute + 15)
	}
}