package saber

import (
	"strconv"
	"strings"

	"github.com/chengxuncc/saber/internal/x"
)

func Int(i interface{}) int {
	if x.IsNil(i) {
		return 0
	}
	switch v := i.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int(v)
	case int8:
		return int(v)
	case uint16:
		return int(v)
	case int16:
		return int(v)
	case uint32:
		return int(v)
	case int32:
		return int(v)
	case uint64:
		return int(v)
	case int64:
		return int(v)
	case int:
		return v
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		return x.Must(strconv.Atoi(strings.TrimSpace(v))).(int)
	case []byte:
		return x.Must(strconv.Atoi(strings.TrimSpace(string(v)))).(int)
	default:
		panic("unknown type to int")
	}
}
