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
