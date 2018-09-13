package debugo

import "strings"

// Match is wildcard matcher
func Match(pattern, str string) bool {
	if pattern == "" {
		return false
	}
	p := strings.Replace(pattern, "*", "", -1)
	if strings.HasPrefix(pattern, "*") && strings.HasSuffix(pattern, "*") {
		return strings.Contains(str, p)
	}
	if strings.HasPrefix(pattern, "*") {
		return strings.HasSuffix(str, p)
	}
	if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix(str, p)
	}
	return pattern == str
}
