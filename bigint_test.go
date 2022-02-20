package bigint

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyByValue(t *testing.T) {
	a := New(10)
	b := a
	a.i.Add(&a.i, big.NewInt(1))
	assert.True(t, a.EQ(New(11)))
	assert.True(t, b.EQ(New(10)))
}

func TestUnmarshalJSON(t *testing.T) {
	expected, _ := FromString("123456785647326594650894650432875047285643207584305843275", 10)
	a := New(0)
	b := a
	err := json.Unmarshal([]byte("123456785647326594650894650432875047285643207584305843275"), &b)
	assert.NoError(t, err)
	assert.True(t, a.EQ(New(0)))
	assert.True(t, b.EQ(expected))
}
