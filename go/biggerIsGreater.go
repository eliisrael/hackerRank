package main

func biggerIsGreater(w string) string {
	// Write your code here
	var i, j int
	for i = len(w) - 1; i > 0; i-- {
		if w[i] > w[i-1] {
			break
		}
	}
	if i == 0 {
		return "no answer"
	}
	for j = len(w) - 1; j >= i; j-- {
		if w[j] > w[i-1] {
			break
		}
	}
	w = swap(w, i-1, j)
	w = reverse(w, i)
	return w
}

func swap(w string, i, j int) string {
	r := []rune(w)
	r[i], r[j] = r[j], r[i]
	return string(r)
}

func reverse(w string, i int) string {
	r := []rune(w)
	for j, k := i, len(r)-1; j < k; j, k = j+1, k-1 {
		r[j], r[k] = r[k], r[j]
	}
	return string(r)
}
