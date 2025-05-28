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

// Calculate computes the optimal number of packs needed to fulfill the item order.
//
// It uses a dynamic programming approach to minimize:
// 1. Extra items sent (i.e., items sent minus items ordered)
// 2. Number of packs used (in case of equal extra)
//
// Returns an error if the input is invalid.
func (p *PackingResult) Calculate(itemsOrdered int, packSizes []int) error {
	// Validate input
	if itemsOrdered <= 0 {
		return ErrInvalidItemCount
	}
	if len(packSizes) == 0 {
		return ErrNoPackSizes
	}

	// Prioritize larger packs first
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	const overshoot = 1000 // Limit how much we can exceed itemsOrdered by
	limit := itemsOrdered + overshoot

	minExtra := math.MaxInt   // Best extra items found so far
	minPacks := math.MaxInt   // Best number of packs used for current minExtra
	best := make(map[int]int) // Best combination of packs found so far

	// Define a state struct to hold intermediate DP results
	type state struct {
		extra int         // Extra items for current state
		packs int         // Total packs used in current state
		used  map[int]int // Map of pack size to quantity used
	}

	dp := make([]*state, limit+1)                     // dp[i] holds best state for total items = i
	dp[0] = &state{packs: 0, used: make(map[int]int)} // Start from zero items

	// Iterate over all reachable totals
	for i := 0; i <= limit; i++ {
		if dp[i] == nil {
			continue // Skip unreachable totals
		}
		for _, size := range packSizes {
			nextTotal := i + size
			if nextTotal > limit {
				continue // Avoid exceeding search limit
			}

			// Copy current pack usage
			newUsed := make(map[int]int)
			for k, v := range dp[i].used {
				newUsed[k] = v
			}
			newUsed[size]++

			newPacks := dp[i].packs + 1
			extra := nextTotal - itemsOrdered

			// Update best result if:
			// - We found fewer extra items
			// - Or we found same extra but fewer packs
			if nextTotal >= itemsOrdered && (extra < minExtra || (extra == minExtra && newPacks < minPacks)) {
				minExtra = extra
				minPacks = newPacks
				best = newUsed
			}

			// Update DP table if this path uses fewer packs
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
