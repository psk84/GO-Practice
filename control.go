package main

import (
	"fmt"
)

/*
	제어구조 (for , if, Switch)
*/

func main() {

	//반복문
	for i := 0; i < 10; i++ {
		//if조건문
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	j := 20

	//switch조건문
	switch j {
	case 10:
		fmt.Println("십")
	case 20:
		fmt.Println("이십")
	case 30:
		fmt.Println("삼십")
	case 40:
		fmt.Println("사십")
	default:
		fmt.Println("알수없는 숫자")
	}
}
