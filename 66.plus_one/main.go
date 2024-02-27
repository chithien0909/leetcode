package main

func plusOne(digits []int) []int {
	return increase(digits)
}

func increase(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	lDigit := digits[len(digits)-1]

	lDigit += 1
	if lDigit > 9 {
		lDigit = 0
		r := increase(digits[:len(digits)-1])
		return append(r, lDigit)
	}
	digits[len(digits)-1] = lDigit
	return digits
}
