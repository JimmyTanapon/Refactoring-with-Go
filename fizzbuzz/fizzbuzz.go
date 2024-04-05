package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
	if isFizzBuzz(n) {
		return "FizzBuzz"
	}
	if isFizz(n) {
		return "Fizz"
	}
	if isBuzz(n) {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func isFizz(num int) bool {
	return num%3 == 0
}
func isBuzz(num int) bool {
	return num%5 == 0
}
func isFizzBuzz(num int) bool {
	return num%15 == 0
}
