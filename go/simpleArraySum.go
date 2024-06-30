package main

func simpleArraySum(ar []int) int {
	// Write your code here
	var sum int = 0
	for _, v := range ar {
		sum += v
	}
	return sum
}
