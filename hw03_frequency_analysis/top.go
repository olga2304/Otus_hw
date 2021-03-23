package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func ExFieldsFunc(s string) []string {
	s = strings.ToLower(s)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(s, f)
}

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
	arr := ExFieldsFunc(str)

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

	sort.Strings(res)

	return res
}
