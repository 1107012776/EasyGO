package tests

import (
	_ "github.com/1107012776/EasyGO/core"
	"github.com/1107012776/EasyGO/server"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Server(t *testing.T) {
	server.Listen("8081", "D:/project/ecc/api/static")
	assert.Equal(t, true, true)
}
