package x

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringsToInterfaces(t *testing.T) {
	rq := require.New(t)
	rq.Equal([]interface{}{}, StringsToInterfaces())
	rq.Equal([]interface{}{"test"}, StringsToInterfaces("test"))
}
