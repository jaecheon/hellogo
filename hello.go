package main

import (
	"fmt"
	"math"
	"os"

	"github.com/jaecheon/greetings"
	"jclee.com/gochannel"
	"jclee.com/gocrawler"
	"jclee.com/goroutine"
)

func main() {
	// fmt.Println("say yaaaahhhh")
	// printEnvs()
	// printSlice()
	// initShape()

	// msg, err := greeting2.Hello("jclee")
	// if(err != nil) {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(msg)

	//1. greeting2 폴더에서 아래 명령실행
	// go mod init jclee.com/greeting2
	//2. 현재 폴더에서 아래 명령실행(git repository없이 로컬에서 사용할 시 사용)
	// go mod edit -replace jclee.com/greeting2=./greeting2
	msg2, err2 := greetings.Hello("jclee")
	if(err2 != nil) {
		fmt.Println(err2.Error())
	}

	fmt.Println(msg2)

	goroutine.Gosay()

	gochannel.Gochannel()

	// 해결할 문제 남음
	// gogeneric.Gogen()

	gocrawler.RunQuery("https://hani.co.kr")
}

func printEnvs() {
	//모든 환경변수 출력
	for index, env := range os.Environ() {
		fmt.Println(index, env)
	}
}

// http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice
func printSlice() {
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스 리터럴값 지정
	a[1] = 1
	fmt.Println(a)
}

// http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map
func printMap() {
	myMap := map[string]string{
		"A":    "Apple",
		"B":    "Banana",
		"C":    "Charlie",
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// Map 키 체크
	// val exists := myMap["MSFT"]
	// if !exists {
	// 	fmt.Println("No MSFT ticker - myMap")
	// }

	// 값이 없으면 nil 혹은 zero 리턴
	noData := myMap["others"]
	fmt.Println(noData)

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

// http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func initShape() {
	r := Rect{10., 20.}
	c := Circle{10}

	showShape(r, c)
}

func showShape(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}
