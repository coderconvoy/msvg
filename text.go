package msvg

import (
	"fmt"
	"unicode/utf8"
)

func Wrap(s string, n int) []string {
	gl := 0 // glyph as k will be char in seq
	lastb := 0
	last_ := 0
	dsh := ""
	res := []string{}
	for k, v := range s {
		switch v {
		case '\n', '\r':
			res = append(res, s[lastb:k]+dsh)
			lastb = k + 1
			last_ = k + 1
			dsh = ""
			gl = 0
			continue

		case ' ':
			last_ = k
			dsh = ""
		case '-':
			last_ = k
			dsh = "-"

		}
		if gl > n && last_ > lastb {
			res = append(res, s[lastb:last_]+dsh)
			lastb = last_ + 1
			fmt.Println("last_,k", lastb, k)
			scut := s[last_:k]
			gl = utf8.RuneCountInString(scut)
			dsh = ""
			continue
		}
		gl++

	}

	res = append(res, s[lastb:])
	return res
}
