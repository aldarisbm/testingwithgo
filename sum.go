package main

func Sum(n []int) int {
	res := 0
	for _, num := range n {
		res += num
	}
	return res
}
