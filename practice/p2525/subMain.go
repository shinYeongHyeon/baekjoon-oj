package p2525

import (
	"bufio"
	"fmt"
	"os"
)

func Exec() {
	var hour int
	var minute int
	var spendMinute int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &hour, &minute)
	fmt.Fscanln(reader, &spendMinute)

	minute += spendMinute

	if minute == 60 {
		hour++
		minute = 0
	} else if minute > 60 {
		hour += minute / 60
		minute %= 60
	}

	if hour >= 24 {
		hour %= 24
	}

	fmt.Println(hour, minute)
}
