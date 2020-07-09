package util

func IsEmpty(v ...interface{}) bool {
	for _, t := range v {
		if t == nil || t == "" || t == "undefined" {
			return true
		}
	}
	return false
}

func IsNotEmpty(v ...interface{}) bool {
	return !IsEmpty(v)
}
