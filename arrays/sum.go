package arrays

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (result []int) {

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (result []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return
}
