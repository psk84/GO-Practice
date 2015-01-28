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
	c3 := Circle1{x: 0, y: 0, r: 5}

	//구조체 초기화 방법4
	c4 := Circle2{0, 0, 5}

	//구조체 필드 접근
	fmt.Println(c4.x, c4.y, c4.r)

	c4.x = 1
	c4.y = 2

	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)
	fmt.Println(c3.x, c3.y, c3.r)

	//메서드 함수 호출
	//이렇게 할경우 가독성이 높아지며 &연산자가 필요하지 않다
	fmt.Println(c4.area())

	//포함관계
	a := new(Android)
	b := new(Android1)

	a.Person.Talk()
	//a.Talk()  => 호출 되지 않는다. Android 구조체는 포함관계가 아니기 때문이다.

	b.Person.Talk()
	b.Talk() // Android1 은 포함관계이기 때문에 Android1에서도 Talk() 호출이 가능하다.

	//인터페이스
	fmt.Println(totalArea(ll))
}

/*
	메서드
	- func 키워드와 함수명 사이에 "수신자" 추가
	- 수신자는 이름과 타입이 있다는 점에서 매개변수와 비슷하지만 이렇게 함수를 생성하면
	- .연산자를 이용하여 해당 함수를 호출 할수 있다.
*/

func (c *Circle2) area() float64 {
	return math.Pi * c.r * c.r
}

/*
	포함관계 - 익명 필드라고도 알려져 있다.
	구조체의 필드는 보통 has-a관계를 나타낸다.
*/

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("안녕, 내 이름은 ", p.Name, "야")
}

//아래의 경우 Android가 Person을 가지고 있다기보다는(Android has a Person)
//Android가 Person이다.(Android is a Person) 이다. 하지만 직접적으로 접근이 어렵다.
type Android struct {
	Person Person
	Model  string
}

//포함타입을 이용해 포함관계를 지원한다.
//익명 필드라고도 알려져있는 포함타입은 아래와 같다.
//타입을 사용하고 이름을 주지 않는다.
//Person타입명을 이용해 구조체에 접근 할수 있다.
type Android1 struct {
	Person
	Model string
}

/*
	인터페이스
	- 서로다른 구조체에서 유사한 함수를 호출하기위한 유사성을 명시적으로 만들어주는 수단을 제공한다.
	- 구조체와 마찬가지로 "type 키워드/이름/interface 키워드"로 구현한다.
	- 하지만 필드를 정의하지 않고 메서드 집합을 정의한다. 메서드 집합은 반드시 구현해야하는 메서드 목록에 해당한다.
*/

type Shape interface {
	getArea() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.getArea()
	}

	return area
}
