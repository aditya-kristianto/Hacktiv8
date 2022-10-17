package helper

func Sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return total
}

func DivideBy2(num int) (total int) {
	if num%2 != 0 {
		return -1
	}

	return num / 2
}

func IsPrime(num int) bool {
	count := 0
	for i := 2; i < num; i++ {
		if num%i == 0 {
			count++
		}
	}

	return count == 0
}
