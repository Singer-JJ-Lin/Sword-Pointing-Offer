package _45_把数组排成最小的树

import (
	"fmt"
	"sort"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", i, j)
		y := fmt.Sprintf("%d%d", i, j)
		return x < y
	})

	var res string
	for _, num := range nums {
		res += fmt.Sprintf("%d", num)
	}
	return res
}
