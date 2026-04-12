package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	//var (
	//	input  = 2
	//	output = 2
	//)
	//
	//actual := AddOne(input)
	//if actual != output {
	//	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	//} else {
	//	t.Logf("AddOne(%d) = %d; want %d", input, actual, output)
	//}

	assert.Equal(t, AddOne(3), 3, "AddOne should be 3")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Executing")
}
