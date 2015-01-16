package main

import "fmt"


/*
	문자열 예제

func main() {

	
	var str1 string= "First Statement"
	str2 := "Second Statement"

	fmt.Println(str1)	//str1 value 출력
	fmt.Println(str2)	//str2 value 출력
	fmt.Println(len(str2))	//str2 value 출력
	fmt.Println(str2[1])	//str2의 두번째 value의 byte값 출력
	fmt.Println(str1 + str2)	//str1 과 str2 value를 합친다.
}
*/

/*
	변수 범위 , 상수 예제 


var str string= "전역 변수"
const CONSTANT_VALUE int = 1000;

func main() {
	fmt.Println("전역 변수 호출(main 함수) : " + str)

	// CONSTANT_VALUE = 2000; // 상수에 갑을 대입하는 처리를 하면 runtime시 'cannot assing to [상수명]'으 에러 발생
	fmt.Println("상수값 호출 : ",CONSTANT_VALUE)
	func1()
}


func func1() {
	fmt.Println("전역 변수 호출(func1 함수) : " + str)
}

*/

/*
	제어구조 (for , if, Switch)


func main() {
	
	//반복문
	for i := 0; i < 10; i++ {
		//if조건문
		if i%2==0 {
			fmt.Println(i)	
		}
	}

	j:=10

	//switch조건문
	switch j (
		case 10: fmt.Println("십")
		case 20: fmt.Println("이십")
		case 30: fmt.Println("삼십")
		default: fmt.Println("나머지요")
	)
}
*/

/*
	배열 , 슬라이스 , 
*/
func main() {
	var x [5]float64	//배열 선언은 다음과 같이 '변수명 [크기]타입' 으로 명시한다.
	x[4] = 100
	fmt.Println(x)

	fmt.Println("기본 배열 출력")
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println("형변환및 숫자 추출(len, 타입())")
	fmt.Println(total / float64(len(x)))	//len()의 반환형은 int, 형변환에서는 함수와 같이 "타입(변환할 변수)" 로 사용하면 된다.

	total = 0

	//아래와 같이 "_" 는 for문에게 사용하지 않는 변수라는것을 지정한다.
	//이유는 go에서의 for는 사용하지 않는 변수를 사용하면 컴파일오류가난다.
	// "range 변수" 는 순차적으로 변수의 내용을 꺼낸다. 자바의 foreach와 비슷한 문법이다.
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	//배열 초기화
	y := [5]int{1,2,3,4,5}

	for _, value := range y{
		fmt.Println(value)
	}

	//배열 초기화 팁

	//아래와 같이 마지막요소에 ,를 적어도 사용에는 무방하며
	//이후 값의 변화에서 주석으로만 처리해주면 되기때문에 편리하다.
	z := [4]int{
		1,
		2,
		3,
		4,
		// 5,
	}

	for _, value := range z{
		fmt.Println(value)
	}

	//슬라이스
	//배열의 일부 
	//배열과의 차이점 : 크기가 가변이다.

	//슬라이스 선언 (크기에 값을 넣지 않아 초기에 0이 슬라이스가 생성된다.)
	var slice1 []int
	fmt.Println("빈 슬라이스 : ", slice1)

	//슬라이스 생성시에는 make함수를 사용하여 생성한다.
	slice2 := make([]int, 5)
	fmt.Println("make 슬라이스 : ", slice2)

	//make함수에서 3번째는 인수는 해당 슬라이스가 점유하는 공간이다.
	slice3 := make([]int, 5, 10)
	fmt.Println("make 슬라이스(3번째 인수10사용): ", slice3, len(slice3))	

}
