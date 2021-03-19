package util

import "regexp"

func FindAllString(s string, exp string) []string {
	reg := regexp.MustCompile(exp)
	return reg.FindAllString(s, -1)
}

func IsMatch(s string, exp string) bool {
	if b, err := regexp.MatchString(exp, s); err != nil {
		return false
	} else {
		return b
	}
}
