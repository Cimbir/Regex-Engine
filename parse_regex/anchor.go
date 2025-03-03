package parseregex

func IsAnchor(c byte) bool {
	return c == '^' || c == '$'
}
