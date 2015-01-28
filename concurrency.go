/*
	Go에서는 동시성을 지원하기 위해 고루틴(goroutine)과 채널(channel)가 존재한다.
*/

package main

import (
	"fmt"
	// "math/rand"
	"time"
)

/*
	고루틴 예제
	- 고루틴은 다른 함수를 동시에 실행 할수 있는 함수
	- 고루틴을 생성하려면 go 키워드 다음에 함수를 호출하면 된다.


	아래는 두개의 고루틴으로 구성되었다.
	- main 과 go f(0)이 호출될때 만들어진다.
	- go로 실행된 함수 다음의 로직은 바로 실행된후 상위 고루틴이 끝나면 해당 함수 실행은 중단된다.
	- 고루틴은 생성에 비용이 많이 들지 않기 때문에 손쉽게 생성 할수 있다.


func f(n int) {
	for i := 0; i < 1000; i++ {
		fmt.Println(n, ":", i)

		//interval을 주어 처리를 확인해본다.
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}
*/

/*
	채널 예제
	- 채널은 두 고루틴이 서로 통신하고 실행흐름을 동기화 하는 수단을 제공
	- 채널의 타입은 chan이라는 키워드 다음에 채널에 전달되는 것의 타입을 지정해서 나타낸다.
	- <- 연산자는 채널에 메시지를 전달하고 채널로부터 메세지를 전달받는데 사용된다.

	채널의 방향
	위의 pinger는 채널에 메세지를 보내는 함수이다.
	이경우 채널을 보내는 채널로 방향을 지정할수 있다.
	"chan<-" 와 같이 chan 키워드 오른쪽에 "<-"왼쪽연산자를 사용하면 보내는 채널로 지정
	반대로 "<-chan" 이면 받는 채널로 방향을 지정할수 있다.

	// func pinger(c chan<- string) {
	// for i := 0; ; i++ {
	// 	c <- "ping" // ping을 전달한다는 의미
	// }

	// func printer(c <-chan string) {
	// for {
	// 	msg := <-c //메세지를 받아 그것을 msg에 저장한다
	// 	fmt.Println(msg)
	// 	time.Sleep(time.Second * 1)
	// }

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // ping을 전달한다는 의미
	}
}

func printer(c chan string) {
	for {
		msg := <-c //메세지를 받아 그것을 msg에 저장한다
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}


func main() {
	var c chan string = make(chan string)

	// 채널 c를 이용하여 두개의 고루틴이 동기화 된다.
	// pinger가 메시지를 채널에 전송하려고 시도할 경우 printer는 해당 메세지를 받을 준비가 될때까지 대기한다.(블로)
	go pinger(c)
	go ponger(c) // ponger를 추가하는 경우 "ping" 과 "pong"이 번갈아 가며 출력된다.
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
*/

/*
	Select
	- switch와 비슷하게 동작하지만 채널에 대해서만 동작하는 select라는 구문이 존재
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			//select를 사용하여 c1채널과 c2에서 전달되는 메세지를 다르게 받아 처리한다.
			select {
			case msg1 := <-c1:
				fmt.Println("select case1 :: ", msg1)
			case msg2 := <-c2:
				fmt.Println("select case2 :: ", msg2)
			//time.After는 채널을 생성한 후 지정한 시간이 지나면 현재시간을 해당 체널로 보낸다.
			case <-time.After(time.Second):
				fmt.Println("timeout")
			//default는 준비된 채널이 없을 경우 즉시 실행된다.
			default:
				fmt.Println("nothing ready ")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
