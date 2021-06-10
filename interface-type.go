package main

import "fmt"

func Map(array []interface{}, callback func(interface{}, int) interface{}) []interface{} {
	var newArray []interface{}
	for i, v := range array {
		newArray = append(newArray, callback(v, i));
	}
	return newArray
}

func main() {
	array := []interface{} { 1, "2", 3, "4" }
	newArray := Map(array, func(each interface{}, index int) interface{} {
		return "wonderful"
	})
	fmt.Println(newArray)
}