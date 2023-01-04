package redisdemo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPing(t *testing.T) {
	rc := GetRedisClient()
	pong, err := rc.GetPing()

	as := assert.New(t)

	as.NotEmpty(pong)
	as.Empty(err)
}
