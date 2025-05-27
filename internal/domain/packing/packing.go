package packing

import (
	"math"
	"sort"
)

type PackingResult struct {
	PacksUsed map[int]int
}

func (p *PackingResult) Calculate(itemsOrdered int, packSizes []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	best := make(map[int]int)
	minExtra := math.MaxInt
	minPacks := math.MaxInt

	var dfs func(idx, total, packs int, used map[int]int)
	dfs = func(idx, total, packs int, used map[int]int) {
		if total >= itemsOrdered {
			extra := total - itemsOrdered
			if extra < minExtra || (extra == minExtra && packs < minPacks) {
				minExtra = extra
				minPacks = packs
				best = make(map[int]int)
				for k, v := range used {
					if v > 0 {
						best[k] = v
					}
				}
			}
			return
		}
		if idx >= len(packSizes) {
			return
		}

		size := packSizes[idx]
		maxCount := (itemsOrdered + size - 1) / size
		for i := 0; i <= maxCount; i++ {
			next := make(map[int]int)
			for k, v := range used {
				next[k] = v
			}
			if i > 0 {
				next[size] = i
			}
			dfs(idx+1, total+i*size, packs+i, next)
		}
	}

	dfs(0, 0, 0, make(map[int]int))
	p.PacksUsed = best
}
