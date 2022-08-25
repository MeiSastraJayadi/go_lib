package lib

func Factorial(n int32) int64 {
	if n < 2 {
		return 1
	}
	return int64(n) * Factorial(n-1)
}
