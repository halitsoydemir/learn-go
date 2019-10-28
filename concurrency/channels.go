package concurrency

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("calculated random values is :{}",value)
	values <- value
}

func ChannelTestDeadlockExceptionThrown(){
	fmt.Println("Test channel starting...")
	values := make(chan int)
	defer close(values)

	go CalculateValue(values)
	//go CalculateValue(values)

	//get val
	//value := <-values
	//fmt.Println(value)

	for val := range values {
		fmt.Println("Consumed created message {}",val)
	}

	//value = <-values
	//fmt.Println(value)
}

func ChannelTest(){
	fmt.Println("Test channel starting...")
	values := make(chan int)
	defer close(values)

	go CalculateValue(values)
	go CalculateValue(values)

	//get val
	value := <-values
	fmt.Println(value)

	value = <-values
	fmt.Println(value)
}

func ChannelSecondaryTest(){
	fmt.Println("Test channel starting...")
	values := make(chan int)

	go CalculateValue(values)
	go CalculateValue(values)

	for {
		select {
		case msg := <- values:
			fmt.Println("Consumed created message {}",msg)
		default:

		}
	}

	close(values)

}

func produce(ch chan int) {
	ch <- rand.Intn(10)
}

func EmresMethod() {

	ch := make(chan int)
	ch2 := make(chan int)

	go produce(ch)
	go produce(ch)
	go produce(ch2)

	end := 0
	for {
		fmt.Println("for")
		select {
		case x := <-ch:
			fmt.Println("value from ch : ", x)
			end++
		case x := <-ch2:
			fmt.Println("value from ch2 : ", x)
			end++
		}

		if end == 2 {
			break
		}
	}

	close(ch)
	close(ch2)
}