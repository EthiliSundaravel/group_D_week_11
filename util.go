package util

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

// ReverseString reverses a string
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// SumOfSlice returns the sum of elements in a slice
func SumOfSlice(slice []int) int {
    sum := 0
    for _, v := range slice {
        sum += v
    }
    return sum
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
    if n < 0 {
        return -1 // Error case for negative input
    }
    if n == 0 || n == 1 {
        return 1
    }
    return n * Factorial(n-1)
}
