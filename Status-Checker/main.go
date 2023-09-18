package main

import (
	"fmt"
	"net/http"
	"time"
)

var websites = []string{
	"http://amazon.com",
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
}

func main() {
	channel := make(chan string)

	for _, website := range websites {
		go checkLink(website, channel)
	}

	for link := range channel {
		go func(childLink string) {
			time.Sleep(5 * time.Second)
			checkLink(childLink, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		channel <- link
		return
	}

	fmt.Println(link, "is up")
	channel <- link
}
