package tests

import (
	"github.com/1107012776/EasyGO/src"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_I_Am(t *testing.T) {
	src.Hello()
	assert.Equal(t, true, true)
}
