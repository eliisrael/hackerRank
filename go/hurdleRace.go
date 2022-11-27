// https://www.hackerrank.com/challenges/the-hurdle-race/problem?isFullScreen=true
package main

import "fmt"

func hurdleRace(ar []int, k int) int {

	// Write your code here
	var greatest int = ar[0]
	for i := 1; i < len(ar); i++ {

		if ar[i] > greatest {
			greatest = ar[i]
		}

	}

	var potionsAmount int = greatest - k
	if potionsAmount < 0 {
		return 0
	}
	return potionsAmount

}

func main() {

	//a := []int{1, 2, 3, 4, 5, 10}
	a := []int{2, 5, 4, 5, 2}
	var capacity int = 7
	var potionsAmount int = hurdleRace(a, capacity)
	fmt.Println(potionsAmount)
}
