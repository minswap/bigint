package bigint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyByValue(t *testing.T) {
	a := New(10)
	b := a
	a_ptr := fmt.Sprintf("%p", &a.i)
	b_ptr := fmt.Sprintf("%p", &b.i)
	assert.NotEqual(t, a_ptr, b_ptr)
}
