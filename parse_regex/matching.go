package parseregex

func IsMatching(c byte) bool {
	return c == '.' || c == '['
}
