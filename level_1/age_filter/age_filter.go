package main

import (
	"fmt"
)

func age_filter(min int, max int, ages[] int) []int{
	result := make([]int, 0)
	
	for _, age := range ages {
		if(age >= min && age <= max){
			result = append(result, age)
		}
	}
	
	return result
}

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(age_filter(2, 3, []int{1, 2, 3, 4, 5}))
}
