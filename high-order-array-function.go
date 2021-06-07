package main

import "fmt"

func Map(array []int, callback func(int, int) int) []int {
	var newArray []int
	for i, v := range array {
		newArray = append(newArray, callback(v, i));
	}
	return newArray
}

func Filter(array[] int, callback func(int, int) bool) []int {
	var newArray []int
	for i, v := range array {
		if callback(v, i) {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func Reduce(array[] int, callback func(int, int, int) int, initialState int) int {
	result := initialState
	for i, v := range array {
		result = callback(result, v, i)
	}
	return result
}

func main() {
	numbers := []int {1, 2 ,3, 4, 5, 6, 7, 8, 9, 10}
	doubledNumbers := Map(numbers, func (e, i int) int {
		return e * 2;
	})
	fmt.Println(doubledNumbers)

	numbersLowerThan5 := Filter(numbers, func(e, i int) bool {
		return e < 5;
	})

	fmt.Println(numbersLowerThan5)

	sum := Reduce(numbers, func(acc, cur, i int) int {
		return acc + cur
	}, 0)

	fmt.Println(sum)

}