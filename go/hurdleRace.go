// https://www.hackerrank.com/challenges/the-hurdle-race/problem?isFullScreen=true
package main

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
