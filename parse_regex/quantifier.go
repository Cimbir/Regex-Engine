package parseregex

func IsQuantifier(c byte) bool {
	return c == '*' || c == '+' || c == '?' || c == '{'
}
