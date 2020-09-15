package syntax

func EnterClub(age int) bool {
	if korAge := age + 2; korAge < 18 {
		return false
	}
	return true
}

func IsPrintable(ch rune) bool {
	switch ch {
	case '\n', '\t', '\r':
		return false
	}
	return false
}
