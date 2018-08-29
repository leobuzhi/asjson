package common

func Isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func Isdigit1to9(c byte) bool {
	return c >= '1' && c <= '9'
}
