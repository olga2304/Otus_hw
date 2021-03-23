package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	arr := strings.FieldsFunc(str, func(r rune) bool {
		return strings.ContainsRune("0123456789\t\n\r ", r)
	})

	counts := make(map[string]int)
	p := make([]string, 0, len(counts))
	res := make([]string, 0, 10)
	t := 10

	for _, a := range arr {
		if _, k := counts[a]; !k {
			p = append(p, a)
		}
		counts[a]++
	}

	sort.SliceStable(p, func(i, j int) bool {
		return counts[p[i]] > counts[p[j]]
	})

	if len(p) < 10 {
		t = len(p)
	}

	res = append(res, p[:t]...)

	return res
}
