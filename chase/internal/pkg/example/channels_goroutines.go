package example

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type urlInfo struct {
	Url     string
	Elapsed float64
}

// Note: This channel is only a writer and can't read from the channels being passed in.
func responseTime(c chan<- urlInfo, stop chan<- bool, url string) {
	start := time.Now()

	if url == "stop" {
		stop <- true
		return
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	c <- urlInfo{
		Url:     url,
		Elapsed: time.Since(start).Seconds(),
	}
}

func BufferedRun() {
	messages := make(chan string, 2)

	messages <- "data"
	messages <- "again"
	// messages <- "wont work as above the buffer size"

	// TODO make this a loop
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// TODO fix with the main function as I think this is largely done
func ChannelSynchronising() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func Sending() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ChannelsRun() {
	// urls := make([]string, 3)
	urls := []string{
		"https://www.australia.gov.au",
		"https://www.google.com",
		"https://www.amazon.com.au",
		// "stop",
	}

	// Buffered channel (up to the max of the number of URLs being processed)
	c := make(chan urlInfo, len(urls))
	// Unbuffered channel
	stop := make(chan bool)

	// TODO document more about the ticker and how to use it
	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()

	for _, u := range urls {
		go responseTime(c, stop, u)
	}

	// Sets up the blocking for channel processes
	count := 0
	for {
		// TODO document about the select operation
		select {
		//TODO fix this to exit so it works consistenly rather than being overridden by the ticker` if they come in at the same time how to order
		// Timeout on the blocking code. I.e. if this section isn't complete after the set time the timeout will be enacted and break the loop.
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout of 50ms has been reached.")
			return
		case t := <-t.C:
			fmt.Printf("Ticker time is %d.\n", t)

		// TODO code this up a bit better so it waits for all work and exits naturally rather than all of these returns
		case <-stop:
			fmt.Printf("Stop received.")
			return

		case data := <-c:
			fmt.Printf("%s took %v seconds \n", data.Url, data.Elapsed)
			// General check as I can't use a for loop check overall as I have other case options rather than just the one for the urls.
			count++
			if count >= len(urls) {
				return
			}
			// default:
			// 	// This attempts a non-blocking message channel
			// 	// TODO more to configure this all up nicely
			// 	fmt.Printf(".")
		}
	}

	// TODO add a range over channels

}

func init() {
	examples := ExampleRuns{
		{"Buffered Run", BufferedRun},
		{"Channels Run", ChannelsRun},
	}
	GetMyExamples().Add("channels", examples.runExamples)
}
