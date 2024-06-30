package main

import "math"

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	var remainders = make([]int32, k)
	for _, v := range s {
		remainders[v%k]++
	}
	var count int32
	if remainders[0] > 0 {
		count++
	}
	for i := int32(1); i <= k/2; i++ {
		if i != k-i {
			count += int32(math.Max(float64(remainders[i]), float64(remainders[k-i])))
		} else {
			count++
		}
	}
	return count
}
