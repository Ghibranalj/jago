package compiler

import (
	"regexp"
)

const (
	FULL        = 0
	MAIN_METHOD = 1
	SINGLETON   = 2
)

// see the match in compiler_test.go
var regex = []string{
	`class\s+.+\s*{[\S\s]+}`, //FULL
	"public\\s+static\\s+void\\s+main\\s{0,}\\(String\\[\\]\\s+.+", //Main Method
	"[\\S\\s]", // singleton
}

func categorize(code string) int {
	cat := 0
	for i, reg := range regex {
		cat = i
		match, _ := regexp.MatchString(reg, code)
		if match {
			break
		}
	}
	return cat
}
