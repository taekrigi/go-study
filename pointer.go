package main

import "fmt"

func double(a int) int {
	return a * 2
}

func print1(a int) {
	fmt.Println(a)
	fmt.Println(a)
}

func print2(a *int) {
	fmt.Println(*a)
	fmt.Println(&a)
}

func print3(a map[string]string) {
	fmt.Println(a["name"])
}

func print4(a *map[string]string) {
	fmt.Println((*a)["name"])
}

func main() {
	double1 := double(2)
	print1(double1)
	print2(&double1)
	print3(map[string]string{"name": "taekrigi"})
	print4(&map[string]string{"name": "taekrigi"})
}