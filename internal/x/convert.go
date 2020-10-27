package x

func StringsToInterfaces(s ...string) []interface{} {
	res := make([]interface{}, 0, len(s))
	for _, ele := range s {
		res = append(res, ele)
	}
	return res
}
