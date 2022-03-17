package arraysNslices

func Sum(numbers[5]int) int {

	sum := 0

	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumUnspecified(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum...[]int)(sums []int)  {
	lengthOfNums := len(numbersToSum)
	sums = make([]int, lengthOfNums)

	for _, numbers := range numbersToSum {
		sums = append(sums, SumUnspecified(numbers))
	}
return sums
}

