package shared

func RemoveAt(s []any, i int) []any {
	if i == 0 {
		return s[1:]
	}
	if i == len(s)-1 {
		return s[:i]
	}
	return append(s[:i], s[i+1:]...)
}
