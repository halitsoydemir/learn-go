package beginner

import (
	"fmt"
	"time"
)

func SecondlyTicker() {
	fmt.Println("Ticker is starting...")
	go BackgroundTicker()
	fmt.Println("Application continue")
	select{}
}

func BackgroundTicker(){
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("tock")
	}
}
