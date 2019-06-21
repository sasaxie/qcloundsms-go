package util

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSha1Sum(t *testing.T) {
	h := Sha1Sum([]byte("1"))
	assert.Equal(t, h, "356a192b7913b04c54574d18c28d46e6395428ab")
}
