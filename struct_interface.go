package main

import (
	"fmt"
	"math"
)

//type이라는 키워드는 새 타입임을 알려준다.
//그다음에 타입의 이름(Circle)이 오고 struct라는 키워드가
//현재 struct 타입을 정의하고 있음을 가리키고 중괄호 안에 필드 목록이 온다.
//각 필드는 이름과 타입을 가진다.
//함수와 마찬가지로 타입이 동일한 필드는 겹쳐서 쓸 수 있다.
type Circle1 struct {
	x float64
	y float64
	r float64
}

type Circle2 struct {
	x, y, r float64
}

func main() {
	//구조체 초기화 방법1
	var c1 Circle1

	//구조체 초기화 방법2
	c2 := new(Circle1)

	//구조체 초기화 방법3
	c3 := Circle1(x: 0, y: 0, r: 5)

	//구조체 초기화 방법4
	c4 := Circle1(0, 0, 5)

	//구조체 필드 접근
	fmt.Println(c4.x, c4.y, c4.r)

	c4.x = 1
	c4.y = 2

	fmt.Println(c4.x, c4.y, c4.r)	
}