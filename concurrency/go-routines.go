package concurrency

import (
	"fmt"
	"time"
)

func AsyncTest() {
	fmt.Println("Goroutine Tutorial")

	go compute(10)
	go compute(10)

	//dont close app
	var input string
	fmt.Scanln(&input)
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func AnonymousGoroutineMethod(){
	go func(value int) {
		for i := 0; i < value; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}	}(10)
	fmt.Scanln()
}
