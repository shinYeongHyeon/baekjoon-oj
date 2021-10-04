package p2588

import "fmt"

func Exec() {
	var threeDigitNumber1, threeDigitNumber2 int
	fmt.Scanf("%d", &threeDigitNumber1)
	fmt.Scanf("%d", &threeDigitNumber2)
	fmt.Println(threeDigitNumber1 * (threeDigitNumber2 % 10))
	fmt.Println(threeDigitNumber1 * ((threeDigitNumber2 / 10) % 10))
	fmt.Println(threeDigitNumber1 * (threeDigitNumber2 / 100))
	fmt.Println(threeDigitNumber1 * threeDigitNumber2)
}
