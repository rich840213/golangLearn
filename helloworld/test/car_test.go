package car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := new("", 100)
	if err != nil {
		t.Fatal("got errors: ", err)
	}

	if c == nil {
		t.Error("car shoud be nil")
	}
}

func TestNewWithAssert(t *testing.T) {
	c, err := new("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = new("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.name)
}
