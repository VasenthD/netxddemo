package main

import (
	
	"fmt"
)



func main(){
	var s string
    fmt.Scan(&s)
	fmt.Println(s)
	ints := []byte(s)
	for i := 2; i<len(ints)-1; i+=2{
		ints[i],ints[i+1]=ints[i+1],ints[i]
    }

	fmt.Println(string(ints))

}