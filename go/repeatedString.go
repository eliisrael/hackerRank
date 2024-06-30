package main

func repeteadString(s string, n int64) int64 {
	var aCount int64
	var aCountInS int64
	for _, v := range s {
		if v == 'a' {
			aCountInS++
		}
	}
	aCount = aCountInS * (n / int64(len(s)))
	for i := int64(0); i < n%int64(len(s)); i++ {
		if s[i] == 'a' {
			aCount++
		}
	}
	return aCount
}
