package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Arr struct {
	Key   string
	Value int
}

type ArrList []Arr

func (p ArrList) Len() int {
	return len(p)
}

func (p ArrList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p ArrList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Top10(str string) []string {
	arr := strings.FieldsFunc(str, func(r rune) bool {
		return strings.ContainsRune("0123456789\t\n\r ", r)
	})

	counts := make(map[string]int)
	for _, word := range arr {
		counts[word]++
	}

	p := make(ArrList, 0, len(counts))
	for k, v := range counts {
		p = append(p, Arr{k, v})
	}

	sort.Sort(sort.Reverse(p))

	res := make([]string, 0, 10)
	for j, pair := range p {
		if j >= 10 {
			break
		}
		res = append(res, pair.Key)
	}

	return res
}
