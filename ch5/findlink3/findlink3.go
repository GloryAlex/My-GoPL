package main

import (
	"GoPl/ch5/links"
	"fmt"
	"log"
	"os"
)

func main() {
	breadthFirst(getPath, os.Args[1:])
}
func breadthFirst(f func(item string) []string, workList []string) {
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func getPath(url string) (ret []string) {
	entries, _ := os.ReadDir(url)
	for _, fil := range entries {
		if fil.IsDir() {
			ret = append(ret, url+"/"+fil.Name())
		}
		fmt.Println(fil)
	}
	return ret
}
