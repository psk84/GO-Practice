package main

import (
	"fmt"
)

/*
	Pointer 예제


//*는 포인터 변수를 "역참조(dereference)"하는 데도 사용된다.
//포인터를 역참조하면 해당 포인터가 가리키는 값에 접근할 수 있다.
//*xPtr = 0이라고 쓰면 "int 값 0을 xPtr가 참조하는 메모리 위치에 저장하라"라고 말하는 셈이다.
//그렇게 하지 않고 xPtr = 0이라고 쓰면 컴파일로 오류가 발생하는데,
//xPtr은 int가 아니라 또 다른 *int만 할당할 수 있는 *int이기 때문이다.
func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5

	//변수의 주소를 구할 때는 & 연산자를 사용한다
	zero(&x)
	fmt.Println(x)
}
*/

/*
	new 예제
	new 를 통해 포인터를 구한다
*/

func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x는 1
}
