package main

import (
	"GoPl/ch5/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	urls, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return urls
}
func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)
	//get handling links
	go func() { workList <- os.Args[1:] }()

	//create 20 goroutines to fetch every links
	for i := 0; i < 20; i++ {
		go func() {
			for seekLink := range unseenLinks {
				foundLinks := crawl(seekLink)
				go func() { workList <- foundLinks }()
			}
		}()
	}
	//main goroutines get links from work list
	//and send them to unseenLinks
	seen := make(map[string]bool)
	for list := range workList {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				unseenLinks <- url
			}
		}
	}
}
