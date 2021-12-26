package goroutine

// https://go-tour-ko.appspot.com/concurrency/1
// goroutine 은 Go 런타임에 의해 관리되는 경량 쓰레드입니다.
// say의 evaluation은 현재의 goroutine에서 일어나고, say 의 실행은 새로운 goroutine에서 일어납니다.
import (
	"fmt"
	"time"
)


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func Gosay() {
	go say("world")
	say("hello")
}