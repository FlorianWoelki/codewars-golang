package main

import "fmt"

// Sloppy solution
func TwoOldestAges(ages []int) [2]int {
	var tempArray [2]int

	for i := 0; i < len(tempArray); i++ {
		max, maxIndex := 0, 0

		for index, number := range ages {
			if number > max {
				max = number
				maxIndex = index
			}
		}

		ages[maxIndex] = ages[len(ages)-1]
		ages = ages[:len(ages)-1]

		if tempArray[0] != 0 {
			if tempArray[0] >= max {
				tempArray[1] = tempArray[0]
				tempArray[0] = max
			}
		} else {
			tempArray[i] = max
		}
	}

	return tempArray
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}
