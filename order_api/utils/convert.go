package utils

func Ptr2Val[T any](p *T, defaultVal ...T) T {
	if p == nil {
		if len(defaultVal) != 0 {
			return defaultVal[0]
		} else {
			var val T
			return val
		}
	}

	return *p
}

func Val2Ptr[T any](val T) *T {
	return &val
}
