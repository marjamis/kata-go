package main

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

func main() {
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

	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()

	for _, u := range urls {
		go responseTime(c, stop, u)
	}

	// Sets up the blocking for channel processes
	count := 0
	for {
		select {
		//TODO fix this to exit so it works consistenly rather than being overridden by the ticker` if they come in at the same time how to order
		// Timeout on the blocking code. I.e. if this section isn't complete after the set time the timeout will be enacted and break the loop.
		case <-time.After(50 * time.Millisecond):
			fmt.Println("Timeout of 50ms has been reached.")
			return
		case t := <-t.C:
			fmt.Printf("Ticker time is %d.\n", t)
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
		}
	}
}
