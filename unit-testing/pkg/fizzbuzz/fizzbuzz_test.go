package fizzbuzz

import "testing"

func TestFizzBuzz_base_case(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "fizzbuzz base case with 3 and 5",
			fizzMultiple: 3,
			buzzMultiple: 5,
			totalNumber:  20,
			expected:     []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func TestFizzBuzz_base_case_reverse(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "reverse of base case 1 should pass",
			fizzMultiple: 5,
			buzzMultiple: 3,
			totalNumber:  20,
			expected:     []string{"1", "2", "Buzz", "4", "Fizz", "Buzz", "7", "8", "Buzz", "Fizz", "11", "Buzz", "13", "14", "FizzBuzz", "16", "17", "Buzz", "19", "Fizz"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func TestFizzBuzz_total_not_sufficient(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "reverse of base case 1 should pass",
			fizzMultiple: 5,
			buzzMultiple: 3,
			totalNumber:  2,
			expected:     []string{"1", "2"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func TestFizzBuzz_negative_fizzValue(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "negative case where total passed is negative",
			fizzMultiple: -5,
			buzzMultiple: 3,
			totalNumber:  20,
			expected:     []string{"1", "2", "Buzz", "4", "Fizz", "Buzz", "7", "8", "Buzz", "Fizz", "11", "Buzz", "13", "14", "FizzBuzz", "16", "17", "Buzz", "19", "Fizz"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func TestFizzBuzz_negative_buzzValue(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "negative case where total passed is negative",
			fizzMultiple: 5,
			buzzMultiple: -3,
			totalNumber:  20,
			expected:     []string{"1", "2", "Buzz", "4", "Fizz", "Buzz", "7", "8", "Buzz", "Fizz", "11", "Buzz", "13", "14", "FizzBuzz", "16", "17", "Buzz", "19", "Fizz"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func TestFizzBuzz_same_values(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "negative case where total passed is negative",
			fizzMultiple: 5,
			buzzMultiple: 5,
			totalNumber:  20,
			expected:     []string{"1", "2", "3", "4", "FizzBuzz", "6", "7", "8", "9", "FizzBuzz", "11", "12", "13", "14", "FizzBuzz", "16", "17", "18", "19", "FizzBuzz"},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

// Test FizzBuzz with zero total number
func TestFizzBuzz_zero_total(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "zero case",
			fizzMultiple: 5,
			buzzMultiple: 3,
			totalNumber:  0,
			expected:     []string{},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

// Test FizzBuzz with zero fizz multiple
func TestFizzBuzz_zero_fizzMultiple(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "zero case fizz multiple",
			fizzMultiple: 0,
			buzzMultiple: 3,
			totalNumber:  20,
			expected:     []string{},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

// Test FizzBuzz with zero fizz multiple
func TestFizzBuzz_zero_buzzMultiple(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "zero case buzz multiple",
			fizzMultiple: 3,
			buzzMultiple: 0,
			totalNumber:  20,
			expected:     []string{},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}


// Test FizzBuzz with zero fizz multiple
func TestFizzBuzz_zero_fizzbuzzMultiple(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "zero case fizz and buzz multiple",
			fizzMultiple: 0,
			buzzMultiple: 0,
			totalNumber:  20,
			expected:     []string{},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

// Test FizzBuzz with a negative number
func TestFizzBuzz_negative_totalNumber(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int64
		buzzMultiple int64
		totalNumber  int64
		expected     []string
	}{
		{
			name:         "negative case where total passed is negative",
			fizzMultiple: 5,
			buzzMultiple: 3,
			totalNumber:  -20,
			expected:     []string{},
		},
	}
	// Run the tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := assertPanicFizzBuzzFunction(t, FizzBuzz, test)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}
			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], v)
				}
			}
		})
	}
}

func assertPanicFizzBuzzFunction(t *testing.T, f func(total, fizzAt, buzzAt int64) []string, test struct {
	name         string
	fizzMultiple int64
	buzzMultiple int64
	totalNumber  int64
	expected     []string
}) []string {
	defer func() {
		if r := recover(); r == nil {
			t.Log("The code did not panic")
			return
		}
		t.Errorf("The code paniced please check edge cases")
	}()
	actual := f(test.totalNumber, test.fizzMultiple, test.buzzMultiple)
	return actual
}
