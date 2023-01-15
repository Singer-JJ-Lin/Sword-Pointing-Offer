package _61_扑克牌中的顺子

import "sort"

// 排序法
func isStraight1(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {

		// 统计大鬼，小鬼数量
		if nums[i] == 0 {
			joker++
			// 判断是否重复
		} else if nums[i] == nums[i+1] {
			return false
		}
	}

	// 判断是否是顺子
	return nums[4]-nums[joker] < 5
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func isStraight2(nums []int) bool {
	// 开数组用于统计是否访问
	visited := make([]bool, 15)

	// 设置最大值最小值
	minValue, maxValue := 1<<31-1, 0
	for _, item := range nums {

		// 大鬼，小鬼，无需统计
		if item == 0 {
			continue
		}

		// 已经出现当前扑克牌，返回false
		if visited[item] {
			return false
		}

		// 更新
		visited[item] = true
		maxValue = max(maxValue, item)
		minValue = min(minValue, item)
	}

	// 返回
	return maxValue-minValue <= 5
}
