package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	expected := 10
	actual := Sum(nums...)

	assert.Equal(t, expected, actual, "ini message")
	fmt.Println("TestSum")
}

func TestDivideBy2Fail(t *testing.T) {
	num := 9

	expected := -1
	actual := DivideBy2(num)
	require.Equal(t, expected, actual, "ini testdivideby2fail")
	fmt.Println("TestDivideBy2Fail")
}

func TestDivideBy2Success(t *testing.T) {
	num := 10

	expected := 5
	actual := DivideBy2(num)
	require.Equal(t, expected, actual)
	fmt.Println("TestDivideBy2Success")
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		actual   bool
		expected bool
	}{
		{
			actual:   IsPrime(5),
			expected: true,
		},
		{
			actual:   IsPrime(13),
			expected: true,
		},
		{
			actual:   IsPrime(25),
			expected: false,
		},
		{
			actual:   IsPrime(43),
			expected: true,
		},
		{
			actual:   IsPrime(55),
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test ke %d", i), func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual)
		})
	}
}
