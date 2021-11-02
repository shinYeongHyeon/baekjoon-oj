package p2562

import "fmt"

func Exec() {
	var number = 0
	var maxNumber = 0
	var maxIndex = 0

	for i := 0; i < 9; i++ {
		fmt.Scanf("%d", &number)

		if number > maxNumber {
			maxNumber = number
			maxIndex = i
		}
	}

	fmt.Println(maxNumber)
	fmt.Println(maxIndex+1)
}