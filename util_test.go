package util

import "testing"

// Unit test for IsPrime using table-driven tests
func TestIsPrime(t *testing.T) {
    tests := []struct {
        input    int
        expected bool
    }{
        {2, true},
        {4, false},
        {7, true},
        {9, false},
        {11, true},
    }

    for _, test := range tests {
        t.Run("IsPrime", func(t *testing.T) {
            result := IsPrime(test.input)
            if result != test.expected {
                t.Errorf("IsPrime(%d) = %v; expected %v", test.input, result, test.expected)
            }
        })
    }
}

// Unit test for ReverseString with edge case
func TestReverseString(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"", ""},
        {"a", "a"},
        {"ab", "ba"},
    }

    for _, test := range tests {
        result := ReverseString(test.input)
        if result != test.expected {
            t.Errorf("ReverseString(%q) = %q; expected %q", test.input, result, test.expected)
        }
    }
}


// Unit test for SumOfSlice with edge case
func TestSumOfSlice(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {[]int{1, 2, 3, 4}, 10},
        {[]int{}, 0},
        {[]int{-1, -2, -3}, -6},
    }

    for _, test := range tests {
        result := SumOfSlice(test.input)
        if result != test.expected {
            t.Errorf("SumOfSlice(%v) = %d; expected %d", test.input, result, test.expected)
        }
    }
}

// Unit test for Factorial using table-driven tests
func TestFactorial(t *testing.T) {
    tests := []struct {
        input    int
        expected int
    }{
        {6, 720},
        {0, 1},
        {1, 1},
        {-6, -1},
        {10,3628800}, // Error case
    }

    for _, test := range tests {
        result := Factorial(test.input)
        if result != test.expected {
            t.Errorf("Factorial(%d) = %d; expected %d", test.input, result, test.expected)
        }
    }
}

// Benchmark test for IsPrime
func BenchmarkIsPrime(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPrime(10007) // Testing a relatively large prime number
    }
}

// Benchmark test for ReverseString
func BenchmarkReverseString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReverseString("hello world")
    }
}

// Benchmark test for SumOfSlice
func BenchmarkSumOfSlice(b *testing.B) {
    slice := make([]int, 1000) // A slice of 1000 integers
    for i := 0; i < b.N; i++ {
        SumOfSlice(slice)
    }
}

// Benchmark test for Factorial
func BenchmarkFactorial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Factorial(10)
    }
}
