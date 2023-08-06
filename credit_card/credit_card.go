// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var cardnumbers string
// 	fmt.Scanln(&cardnumbers)
// 	cardnumbers = strings.ReplaceAll(cardnumbers, "-", "")
// 	fmt.Println(cardnumbers)

//	if len(cardnumbers)!=16||cardnumbers[0]!='4' || cardnumbers[0]!='5' || cardnumbers[0]!='6'{
//		panic("Invalid")
//	}
//	for _, n := range cardnumbers {
//		if n<'0'||n>'9'{
//			//fmt.Println("Invalid")
//			panic("Invalid")
//		}
//	}
//
// fmt.Println("valid card")
// }
package main

import (
	"fmt"
	"strings"
)

func main() {
	var cardnumbers string
	fmt.Scanln(&cardnumbers)
	cardnumbers = strings.ReplaceAll(cardnumbers, "-", "")
	fmt.Println(cardnumbers)

	if len(cardnumbers) > 16 || len(cardnumbers) < 13 || (cardnumbers[0] != '4' && cardnumbers[0] != '5' && cardnumbers[0] != '3') {
		fmt.Println("Invalid1")
		return
	}
	for _, n := range cardnumbers {
		if n < '0' || n > '9' {
			fmt.Println("Invalid2")
			return

		}
	}
	//fmt.Println("Valid credit card numbers")

	var odd int = 0
	var even int = 0
	var s1, s2 int
	if len(cardnumbers)%2 == 0 {
		s1 = 1
	} else {
		s2 = 1
	}
	for i := s1; i < len(cardnumbers); i += 2 {
		x := ((int(cardnumbers[i])) - '0')
		fmt.Println("odd  : ", x)
		odd = odd + x
	}
	for i := s2; i < len(cardnumbers); i += 2 {
		x := ((int(cardnumbers[i])) - '0')
		fmt.Println("eve  : ", x)
		if x > 4 {
			n := x * 2
			f := n % 10
			l := n / 10
			n = f + l

			even = even + n

		} else {
			even = even + (x * 2)
		}
	}
	fmt.Println("sum1 : ", even)
	fmt.Println("sum2 : ", odd)
	if (odd+even)%10 != 0 {
		fmt.Println("Invalid3")
		return
	}
	fmt.Println("Valid credit card numbers")

}
