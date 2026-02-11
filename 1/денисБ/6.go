package main

import "fmt"

func main(){
	post := [][]string {
		{"go","backend"}, {"git", "go", "tools"},
	}
	uniqueTags := make(map[string]bool)
	for _, tags := range post {
		for _, tag := range tags{
uniqueTags[tag] = true
		}
	}
	for value, tag := range uniqueTags {
		fmt.Println(tag, value)
	}
}