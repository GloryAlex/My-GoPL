package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math2"},
	"databases":             {"data structures"},
	"discrete math2":         {"intro to programming"},
	"formal languages":      {"discrete math2"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	ret := topSort(prereqs)
	for i, s := range ret {
		fmt.Printf("%02d\t%s\n", i, s)
	}
}
func topSort(pres map[string][]string) []string {
	var ret []string
	var seen = make(map[string]bool)
	var visit func(string)
	visit = func(course string) {
		if seen[course] == false {
			for _, preCourse := range pres[course] {
				visit(preCourse)
			}
			seen[course] = true
			ret = append(ret, course)
		}
	}
	for s, _ := range pres {
		visit(s)
	}
	return ret
}
