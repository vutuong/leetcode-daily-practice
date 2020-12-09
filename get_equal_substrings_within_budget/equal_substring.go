package main

import "fmt"

// func abs(number1, number2 int) int {
// 	if number1 <= number2 {
// 		return number2 - number1
// 	}
// 	return number1 - number2
// }

func equalSubstring(s string, t string, maxCost int) int {
	costs := []int{}
	if len(s) == 0 {
		return 0
	}
	for i := range s {
		if int(s[i])-int(t[i]) >= 0 {
			costs = append(costs, int(s[i])-int(t[i]))
		} else {
			costs = append(costs, int(t[i])-int(s[i]))
		}
	}
	fmt.Println(costs)
	maxLenght := 0
	for i := range s {
		lenght := 0
		costBudget := maxCost
		for j := i; j < len(s); j++ {
			if costBudget-costs[j] < 0 {
				break
			}
			lenght = lenght + 1
			costBudget = costBudget - costs[j]
		}
		if lenght > maxLenght {
			maxLenght = lenght
		}

	}
	// maxLenght := 0
	// for i := range s {
	// 	lenght := 0
	// 	costBudget := maxCost
	// 	for j := i; j < len(s); j++ {
	// 		if costBudget-abs(int(s[j]), int(t[j])) < 0 {
	// 			break
	// 		}
	// 		lenght = lenght + 1
	// 		costBudget = costBudget - abs(int(s[j]), int(t[j]))
	// 	}
	// 	if lenght > maxLenght {
	// 		maxLenght = lenght
	// 	}

	// }
	return maxLenght
}

func main() {
	fmt.Println(equalSubstring("???", "cdef", 3))
}
