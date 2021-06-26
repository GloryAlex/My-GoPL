package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(formatUrl(url), channel)
	}
	for range os.Args[1:] {
		fmt.Println(<-channel)
	}
	fmt.Printf("--------\n%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("While reading %s:%v", url, err)
		return
	}
	timeCost := time.Since(start).Seconds()
	_ = resp.Body.Close() //释放内存
	ch <- fmt.Sprintf("%s: timeCost:%.2fs | httpStatus:%s | receiveSize:%7d", url, timeCost, resp.Status, nBytes)
}

func formatUrl(url string) string {
	if strings.HasPrefix(url, "http://") == false {
		return "http://" + url
	} else {
		return url
	}
}
