package main

import (
	"fmt"
)

/*
	변수 범위 , 상수 예제
*/

var str string = "전역 변수"

const CONSTANT_VALUE int = 1000

func main() {
	fmt.Println("전역 변수 호출(main 함수) : " + str)

	// CONSTANT_VALUE = 2000; // 상수에 갑을 대입하는 처리를 하면 runtime시 'cannot assing to [상수명]'으 에러 발생
	fmt.Println("상수값 호출 : ", CONSTANT_VALUE)
	func1()
}

func func1() {
	fmt.Println("전역 변수 호출(func1 함수) : " + str)
}
