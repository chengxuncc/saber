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

func TestInt(t *testing.T) {
	rq := require.New(t)
	Int := func(i interface{}) int {
		v, err := Int(i)
		rq.NoError(err)
		return v
	}
	rq.Equal(0, Int(nil))
	rq.Equal(0, Int(error(nil)))
	rq.Equal(0, Int([]byte(nil)))
	rq.Equal(0, Int(0))
	rq.Equal(1, Int(1))
	rq.Equal(0, Int(int8(0)))
	rq.Equal(1, Int(int8(1)))
	rq.Equal(0, Int(uint8(0)))
	rq.Equal(1, Int(uint8(1)))
	rq.Equal(0, Int(int16(0)))
	rq.Equal(1, Int(int16(1)))
	rq.Equal(0, Int(uint16(0)))
	rq.Equal(1, Int(uint16(1)))
	rq.Equal(0, Int(int32(0)))
	rq.Equal(1, Int(int32(1)))
	rq.Equal(0, Int(uint32(0)))
	rq.Equal(1, Int(uint32(1)))
	rq.Equal(0, Int(int64(0)))
	rq.Equal(1, Int(int64(1)))
	rq.Equal(0, Int(uint64(0)))
	rq.Equal(1, Int(uint64(1)))
	rq.Equal(0, Int(float32(0)))
	rq.Equal(1, Int(float32(1)))
	rq.Equal(0, Int(float64(0)))
	rq.Equal(1, Int(float64(1)))
}
