package main

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12)

}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for i, e := range bill {
		if int32(i) != k {
			sum += e
		}
	}
	sum /= 2
	if sum == b {
		println("Bon Appetit")
	} else {
		println(b - sum)
	}
}
