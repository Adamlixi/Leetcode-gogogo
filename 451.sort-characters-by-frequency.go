package leetcode

import "sort"

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start

type Pair struct {
	Key   byte
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func frequencySort(s string) string {
	sMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = 0
		}
		sMap[s[i]]++
	}
	pl := make(PairList, 0)
	for k, v := range sMap {
		pl = append(pl, Pair{k, v})
	}
	sort.Sort(sort.Reverse(pl))
	var ans []byte
	for _, p := range pl {
		t := p.Value
		k := p.Key
		for i := 0; i < t; i++ {
			ans = append(ans, k)
		}
	}
	return string(ans)
}

// @lc code=end
