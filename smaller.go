package main

import "fmt"

func main() {
	fmt.Println("No. of digits")
	var num int
	fmt.Scanln(&num)

	input := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("enter digits")
		var digits int
		fmt.Scanln(&digits)
		input[i] = digits
	}
	fmt.Println("input: ", input)
	result := smallerNumbersThanCurrent(input)
	fmt.Println("reult: ", result)

}

func smallerNumbersThanCurrent(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {

		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
		result = append(result, count)

	}

	return result
}
