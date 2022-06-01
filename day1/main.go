package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		p := l + (r-1)/2
		if nums[p] == target {
			return p
		} else if nums[p] < target {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return -1
}
