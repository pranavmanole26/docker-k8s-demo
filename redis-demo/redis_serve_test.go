package redisdemo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MRedisConfig struct {
}

func (rc *MRedisConfig) GetPing() (string, error) {
	return "Pong", nil
}

func GetMockRedisClient() *MRedisConfig {
	return &MRedisConfig{}
}

func TestGetPing(t *testing.T) {
	rc := GetMockRedisClient()

	t.Run("Get Ping testing", func(t *testing.T) {
		pong, err := rc.GetPing()
		as := assert.New(t)
		as.NotEmpty(pong)
		as.Equal("pong", strings.ToLower(pong))
		as.Empty(err)
	})

}
