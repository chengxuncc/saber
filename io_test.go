package saber

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	rq := require.New(t)
	rq.Equal("\n", Echo("").Output())
	rq.Equal("test echo\n", Echo("test", "echo").Output())
}
