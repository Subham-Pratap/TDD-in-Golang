package main

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbers ...[]int) []int {

	var sums []int

	for _, val := range numbers {
		sums = append(sums, Sum(val))
	}
	return sums

}
