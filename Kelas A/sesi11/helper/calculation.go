package helper

func Sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}

	return total
}

func Multiply(num1, num2 int) (total int) {
	total = num1 * num2
	return total
}
