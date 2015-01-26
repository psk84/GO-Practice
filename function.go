package main

import (
	"fmt"
	// "os"
)

/*
  function 생성

  정의 방식 :
  함수 이름 (매개변수명, 매개변수 타입,.....) 반환 타입

func average(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
*/

/*
  스택상에 만들어진 함수 호출


func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}
*/

/*
	반환 형의 이름을 명시 해주는 예제
	반환 타입 정의시 (반환 타입 명 반환타입)으로 미리 반환 타입명을 지정할 수 있다.

func main() {
	fmt.Println(f1())
}

func f1() (r int) {
	r = 10
	return
}
*/

/*
	다중값 반환 예제
	반환형 정의시 (반환 타입, 반환타입 , ....)을 통해 다중값 반환을 할수 있다.

func main() {
	x, y := f()
	fmt.Println(x + y)
}

func f() (string, string) {
	return "박", "승규"
}
*/

/*
	가변 함수 예제

	매개변수 타입 지정시 앞에 '...'을 사용하여 0개 이상의 매개변수를 받는다고 지정한다.


func main() {
	fmt.Println(add(1, 2, 3))

	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total / int(len(args))
}
*/

/*
	클로저 예제
	함수 안에 함수를 만드는 것이 가능하다.
	이경우 정의된 클로저에서는 정의된 함수에 있는 지역 변수에도 접근 가능하다.

func main() {
	//변수 정의
	args := []int{1, 2, 3}
	add := func() int {
		total := 0

		//클로저 내부에서 args접근하여 로직 처리
		for _, v := range args {
			total += v
		}

		return total
	}

	fmt.Println("Closer :: ", add())
}
*/

/*
	지연
	- Go에서는 defer라고 하는 특별한 구문을 사용하여 해당 함수가 실행을 완료했을때 실행을 위해 호출 스케쥴러를 잡는다.


func main() {
	//defer를 사용하여 f2()에대한 호출을 해당 함수의 끝으로 옮겼다(지연)
	defer f2()
	f1()

	//가장 좋은 예로는 파일 처리시 파일을 open하여 로직을 처리한후 open한 객체를 반드시 close해줘야 할때 사용하면 좋다.
	// f, _ := os.Open(filename)
	// defer f.Close()

}

func f1() {
	fmt.Println("Delay :: first")
}

func f2() {
	fmt.Println("Delay :: second")
}
*/

/*
	패닉, 복구 예제
	- panic()을 통하여 런타임 오류를 발생할수 있다.
	- recover()를 통하여 란타임 패닉을 처리할 수 있다.
	- receover는 패닉을 중단하고 panicㅇ 호출에 전달된 값을 반환 한다.

func main() {
	defer func() {
		str := recover()
		fmt.Println("Pacnic Value :: ", str)
	}()

	panic("PANIC")
}
*/
