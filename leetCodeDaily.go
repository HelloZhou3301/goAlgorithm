package main

import "slices"

// numRescueBoats
// 2024-05-04
// https://leetcode.com/problems/boats-to-save-people/description/?envType=daily-question&envId=2024-05-04
func numRescueBoats(people []int, limit int) int {
	res := 0
	// 把people从小到大排序
	slices.Sort(people)
	rightCur := len(people) - 1
	// 待运送的人在[leftCur,rightCur]中
	// 每次必运走剩余人中最重的
	// 再尝试带走一个可带走中最重的
	for leftCur := 0; rightCur >= leftCur; rightCur-- {
		// 只剩一人
		if rightCur == leftCur {
			res++
			break
		}
		// 最瘦的可以和最重的一起，把最瘦的一带走
		if people[leftCur]+people[rightCur] <= limit {
			leftCur++
		}
		res++
	}
	return res
}
