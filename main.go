package main

import "fmt"

func main(){
	// var scores []int = []int{50, 75, 66, 20, 32,90,}
	var scores []int = []int{66, 75, 50, 20, 32,90,}
	// var scores []int = []int{90, 75, 50, 20, 32,66,}
	var index int = 0
	var target int = 66

	for i, value := range scores{
		if value == target {
			index = i
		}
	}

	scores = append(scores[:index+1], append([]int {88,}, scores[index+1:]...)...)
	

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}