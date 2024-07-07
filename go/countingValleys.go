package main

func countingValleys(n int32, s string) int32 {
	// Write your code here
	var count, level int32
	for _, v := range s {
		if v == 'U' {
			level++
		} else {
			level--
		}
		if level == 0 && v == 'U' {
			count++
		}
	}
	return count
}
