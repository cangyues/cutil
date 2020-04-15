package util

func IsEmpty(v interface{}) bool {
	if v == nil || v == "" || v == "undefined" {
		return true
	}
	return false
}

func IsNotEmpty(v interface{}) bool {
	return !IsEmpty(v)
}
