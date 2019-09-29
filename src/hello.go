package main

import "fmt"

const helloPrefix = "hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("hj"))
	var a = [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
