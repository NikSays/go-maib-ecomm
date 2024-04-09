package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseDay(t *testing.T) {
	req := CloseDay{}
	err := req.Validate()
	assert.Nil(t, err)
	val, err := req.Values()
	assert.Nil(t, err)
	assert.Equal(t, "command=b", val.Encode())
}
