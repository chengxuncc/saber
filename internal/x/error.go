package x

func Must(i ...interface{}) interface{} {
	for _, v := range i {
		if v == nil {
			continue
		}
		if err, ok := v.(error); ok && err != nil {
			panic(err.Error())
		}
	}
	for _, v := range i {
		return v
	}
	return nil
}
