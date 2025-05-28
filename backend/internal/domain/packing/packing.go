package packing

import (
	"errors"
	"math"
	"sort"
)

var (
	ErrInvalidItemCount = errors.New("number of items must be greater than zero")
	ErrNoPackSizes      = errors.New("at least one pack size must be provided")
)

type PackingResult struct {
	PacksUsed map[int]int
}

// Calculate determines the optimal number of packs to fulfill the item order
// using dynamic programming to minimize extra items and number of packs.
func (p *PackingResult) Calculate(itemsOrdered int, packSizes []int) error {
	if itemsOrdered <= 0 {
		return ErrInvalidItemCount
	}
	if len(packSizes) == 0 {
		return ErrNoPackSizes
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	const overshoot = 1000 // margin above itemsOrdered
	limit := itemsOrdered + overshoot
	minExtra := math.MaxInt
	minPacks := math.MaxInt
	best := make(map[int]int)

	type state struct {
		extra int
		packs int
		used  map[int]int
	}
	dp := make([]*state, limit+1)
	dp[0] = &state{extra: 0, packs: 0, used: make(map[int]int)}

	for i := 0; i <= limit; i++ {
		if dp[i] == nil {
			continue
		}
		for _, size := range packSizes {
			nextTotal := i + size
			if nextTotal > limit {
				continue
			}
			newUsed := make(map[int]int)
			for k, v := range dp[i].used {
				newUsed[k] = v
			}
			newUsed[size]++
			newPacks := dp[i].packs + 1
			extra := nextTotal - itemsOrdered

			if nextTotal >= itemsOrdered && (extra < minExtra || (extra == minExtra && newPacks < minPacks)) {
				minExtra = extra
				minPacks = newPacks
				best = newUsed
			}

			if dp[nextTotal] == nil || newPacks < dp[nextTotal].packs {
				dp[nextTotal] = &state{
					extra: extra,
					packs: newPacks,
					used:  newUsed,
				}
			}
		}
	}

	p.PacksUsed = best
	return nil
}
