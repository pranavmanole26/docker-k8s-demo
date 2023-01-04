package redisdemo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPing(t *testing.T) {
	rc := GetRedisClient()

	t.Run("Get Ping testing", func(t *testing.T) {
		pong, err := rc.GetPing()
		as := assert.New(t)
		as.NotEmpty(pong)
		as.Equal("pong", strings.ToLower(pong))
		as.Empty(err)
	})

}
