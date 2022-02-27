package bigint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.True(t, Min(New(20), New(10), New(10)).EQ(New(10)))
}

func TestMax(t *testing.T) {
	assert.True(t, Max(New(10), New(30), New(20)).EQ(New(30)))
}

func TestSum(t *testing.T) {
	assert.True(t, Sum().EQ(New(0)))
	assert.True(t, Sum(New(5), New(6), New(7)).EQ((New(5 + 6 + 7))))
}

func TestProduct(t *testing.T) {
	assert.True(t, Product().EQ(New(1)))
	assert.True(t, Product(New(5), New(6), New(7)).EQ((New(5 * 6 * 7))))
}
