package msvg

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

		default:
		}
		if gl > n && last_ > lastb {
			res = append(res, s[lastb:last_]+dsh)
			gl = 0
			lastb = last_ + 1
			continue
		}
		gl++

	}

	res = append(res, s[lastb:])
	return res
}
