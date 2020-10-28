package x

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func StringsToInterfaces(s ...string) []interface{} {
	res := make([]interface{}, 0, len(s))
	for _, ele := range s {
		res = append(res, ele)
	}
	return res
}

func Int(i interface{}) (int, error) {
	if IsNil(i) {
		return 0, nil
	}
	switch v := i.(type) {
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case uint8:
		return int(v), nil
	case int8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case int16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case int32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case int64:
		return int(v), nil
	case int:
		return v, nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return Must(strconv.Atoi(strings.TrimSpace(v))).(int), nil
	case []byte:
		return Must(strconv.Atoi(strings.TrimSpace(string(v)))).(int), nil
	default:
		return 0, fmt.Errorf("unsupported converting type %s to int", reflect.TypeOf(i).String())
	}
}
