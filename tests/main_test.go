package tests

import (
	"github.com/1107012776/EasyGO/core"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_I_Am(t *testing.T) {
	core.Hello()
	assert.Equal(t, true, true)
}
