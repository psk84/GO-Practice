package main

import "fmt"

/*
	문자열 예제
*/
func main() {

	var str1 string = "First Statement"
	str2 := "Second Statement"

	fmt.Println(str1)        //str1 value 출력
	fmt.Println(str2)        //str2 value 출력
	fmt.Println(len(str2))   //str2 value 출력
	fmt.Println(str2[1])     //str2의 두번째 value의 byte값 출력
	fmt.Println(str1 + str2) //str1 과 str2 value를 합친다.
}
