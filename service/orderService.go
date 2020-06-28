//
// Copyright (C) 2020 - -
//

package service

import "sort"

// Pair struct
type (
	Pair struct {
		Key   string
		Value int
	}

	// PairList struct
	PairList []Pair
)

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// orderMap : function to sort map by value
func orderMap(noble *map[string]float64) PairList {
	p := make(PairList, len(*noble))
	i := 0
	for k, v := range *noble {
		p[i] = Pair{k, int(v)}
		i++
	}
	sort.Sort(p)
	return p
}
