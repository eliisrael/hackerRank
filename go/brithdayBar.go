package main

func birthdayBar(s []int32, d int32, m int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i <= len(s)-int(m); i++ {
		var sum int32
		for j := i; j < i+int(m); j++ {
			sum += s[j]
		}
		if sum == d {
			count++
		}
	}
	return count
}
