package main

import "fmt"

type Sports interface {
	sound()
}

type Football struct {}

func (football *Football) sound() {
	fmt.Println("pung...! pung...!")
}

type BaseBall struct {}

func (baseball *BaseBall) sound() {
	fmt.Println("ddang...! ddang...!")
}

func PrintSports(sports Sports) {
	sports.sound()
}

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

	football := Football{}
	baseball := BaseBall{}

	PrintSports(&football)
	PrintSports(&baseball)
}