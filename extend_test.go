package baseutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinBytes(t *testing.T) {
	a := "Tom"
	b := "Jerry"
	c := "Puppy"
	sep := []byte("")
	expected := []byte(a + b + c)
	actual := JoinBytes(sep, []byte(a), []byte(b), []byte(c))
	assert.Equal(t, expected, actual)

	sepStr := ";"
	expected = []byte(a + sepStr + b + sepStr + c)
	sep = []byte(sepStr)
	actual = JoinBytes(sep, []byte(a), []byte(b), []byte(c))
	assert.Equal(t, expected, actual)
}
