package example

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func buffered() {
	messages := make(chan string, 2)

	messages <- "data"
	messages <- "again"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronising() {
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

func passingMessages() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

type urlInfo struct {
	URL     string
	Elapsed float64
}

// Note: This channel is only a writer and can't read from the channels being passed in.
func responseTime(c chan<- urlInfo, url string) {
	start := time.Now()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	c <- urlInfo{
		URL:     url,
		Elapsed: time.Since(start).Seconds(),
	}
}

func urlSpeedTesting() {
	urls := []string{
		"https://www.australia.gov.au",
		"https://www.google.com",
		"https://www.amazon.com.au",
	}

	// Buffered channel (up to the max of the number of URLs being processed)
	// This means only this number of messages can be stored waiting for processing
	c := make(chan urlInfo, len(urls))

	// Buffered channel that will only accept on bool value
	// This is used to determine if the loop should stop. As soon as any part of the code sends data on this channel the loop will stop
	stop := make(chan bool, 1)

	// Creates a ticker which can be used for how often the below for loop goes to the next select operation
	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()

	for _, u := range urls {
		go responseTime(c, u)
	}

	// Sets up the blocking for channel processes
	count := 0
	for {
		select {
		// If the stop channel is activated, in this case by the timeout or by completing all the required ping, the loop exits
		case <-stop:
			fmt.Printf("Stop received.")
			return

		// Timeout on the blocking code. I.e. if this section isn't complete after the set time the timeout will be enacted and break the loop
		case <-time.After(60 * time.Millisecond):
			fmt.Println("Timeout of 60ms has been reached.")
			stop <- true

		// Prints out the ticker time this also controls the loop speed before testing the case again
		case t := <-t.C:
			fmt.Printf("Ticker time is %+v.\n", t)

		// If the data channel receives data it prints the details and if the max amount of enteries are delivered the stop channel is sent the message to close
		case data := <-c:
			fmt.Printf("%s took %v seconds \n", data.URL, data.Elapsed)
			// General check as I can't use a for loop check overall as I have other case options rather than just the one for the urls.
			count++
			if count >= len(urls) {
				fmt.Println("All URL's have responded. Exiting...")
				stop <- true
			}
		}
	}
}

func unbuffered() {
	messages := make(chan string)

	go func(messages chan string) {
		fmt.Println(<-messages)
	}(messages)
	messages <- "data"

	go func(messages chan string) {
		fmt.Println(<-messages)
	}(messages)
	messages <- "again"

	// As the messages channel is unbuffered it needs a receiver actively listening otherwise it halts on trying to send
	// data to messages. Hence if you uncomment the below it will stall the application but the above works as each time there is
	// a go func() listening for the input when the string is sent over the messages channel.
	// messages <- "but not now"
}

func init() {
	category := GetCategories().AddCategory("channels")

	category.AddExample("buffered",
		CategoryExample{
			Description: "Basic buffered channel example",
			Function:    buffered,
		})
	category.AddExample("unbuffered",
		CategoryExample{
			Description: "Basic unbuffered channel example",
			Function:    unbuffered,
		})
	category.AddExample("passingMessages",
		CategoryExample{
			Description: "Passing messages between functions",
			Function:    passingMessages,
		})
	category.AddExample("syncing",
		CategoryExample{
			Description: "Channel synchronising i.e. waiting to exit",
			Function:    channelSynchronising,
		})
	category.AddExample("urls",
		CategoryExample{
			Description: "URL Speed Testing",
			Function:    urlSpeedTesting,
		})
}
