package main

func betweenTwoSets(a []int32, b []int32) int32 {
	// Write your code here
	var count int32
	for i := a[len(a)-1]; i <= b[0]; i++ {
		var flag bool
		for _, v := range a {
			if i%v != 0 {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		for _, v := range b {
			if v%i != 0 {
				flag = true
				break
			}
		}
		if !flag {
			count++
		}
	}
	return count
}
