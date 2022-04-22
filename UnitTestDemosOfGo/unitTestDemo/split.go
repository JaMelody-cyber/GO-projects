package split

import "strings"

func Split(s, rem string) (res []string) {
	i := strings.Index(s, rem)
	for i > -1 {
		res = append(res, s[:i])
		s = s[i+1:]
		i = strings.Index(s, rem)
	}
	res = append(res, s)
	return res
}
