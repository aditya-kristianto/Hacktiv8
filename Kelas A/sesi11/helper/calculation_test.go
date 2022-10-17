package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []int{10, 20, 30}

	actual := Sum(nums...)
	expected := 70

	assert.Equal(t, expected, actual)
	// if actual != expected {
	// 	t.Fatalf("expected : %d, got : %d", expected, actual)
	// }

	fmt.Println("unit test sum done...")
}

func TestMultiply(t *testing.T) {
	actual := Multiply(10, 20)
	expected := 80

	require.Equal(t, expected, actual)
	// if actual != expected {
	// 	t.Fail()
	// 	// t.Fail("expected : %d, got : %d", expected, actual)
	// }
	fmt.Println("unit test multiply done...")

}
