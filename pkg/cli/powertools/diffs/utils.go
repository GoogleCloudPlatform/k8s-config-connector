package diffs

import (
	"cmp"
	"strings"
)

func compareStringSlices(l []string, r []string) int {
	lLen := len(l)
	rLen := len(r)
	minLen := min(lLen, rLen)
	for i := 0; i < minLen; i++ {
		lv := l[i]
		rv := r[i]
		if x := strings.Compare(lv, rv); x != 0 {
			return x
		}
	}
	if lLen < rLen {
		return -1
	}
	if rLen > lLen {
		return 1
	}
	return 0
}

func min[T cmp.Ordered](l, r T) T {
	if l < r {
		return l
	}
	return r
}
