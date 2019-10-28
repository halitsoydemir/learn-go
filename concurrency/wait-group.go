package concurrency

import (
	"fmt"
	"sync"
)

func sampleGoRoutine(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}

func TestWithoutWaitGroup() {
	fmt.Println("Hello World")
	go sampleGoRoutine(nil)
	fmt.Println("Finished Execution")
}

func TestWithWaitGroup() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go sampleGoRoutine(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}

func TestWithAnonymousWaitGroup() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go func(url string) {
		fmt.Println("Inside my goroutine ",url)
		waitgroup.Done()
	}("hello.com")
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}