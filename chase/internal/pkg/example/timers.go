package example

// Example from: https://gobyexample.com/timers

import (
	"fmt"
	"time"
)

func timer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

func init() {
	examples := runs{
		{"Basic timer that is set and fires and an additional timer which is stopped", timer},
	}
	GetMyExamples().Add("timers", examples.runExamples)
}
