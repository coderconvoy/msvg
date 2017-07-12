package msvg

func Wrap(s string, n int) []string {
	gl := 0 // glyph as k will be char in seq
	lastb := 0
	last_ := 0
	res := []string{}
	for k, v := range s {
		switch v {
		case '\n', '\r':
			res = append(res, s[lastb:k])
			lastb = k + 1
			last_ = k + 1

		case ' ', '-':
			last_ = k
		default:
		}
		if gl > n && last_ > lastb {
			res = append(res, s[lastb:last_])
			gl = 0
			lastb = last_ + 1
			continue
		}
		gl++

	}

	res = append(res, s[lastb:])
	return res
}
