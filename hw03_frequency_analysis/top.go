package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	arr := strings.Fields(str)

	counts := make(map[string]int)
	p := make([]string, 0)
	t := 10

	for _, a := range arr {
		if _, ok := counts[a]; !ok {
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

	res := make([]string, 0, 10)
	res = append(res, p[:t]...)

	return res
}
