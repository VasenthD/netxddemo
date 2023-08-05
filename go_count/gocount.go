package main

import "fmt"

func main() {
	var count = 0
	var str string
	fmt.Scan(&str)
	element := []byte(str)
	fmt.Println(str)
	for e, _ := range str {
		if string(element[e]) == "g" && string(element[e+1]) == "o" {
			count++
		}
	}
	fmt.Println(count)
}
