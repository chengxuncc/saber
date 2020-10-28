package x

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNil(t *testing.T) {
	var err error
	require.Equal(t, true, IsNil(err))
}
