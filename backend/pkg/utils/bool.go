package utils

// BoolToInt は bool 値を int に変換する (true → 1, false → 0)
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
