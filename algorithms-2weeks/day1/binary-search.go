//配列全体を探索の範囲とします。
//探索の範囲内の中央の要素を調べます。
//目的のキーと中央の要素のキーが一致すれば探索を終了します。
//目的のキーが中央の要素のキーよりも小さければ前半部分を、大きければ後半部分を探索の範囲として 2. へ戻ります。

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {

		pivot := (left + right) / 2

		if nums[pivot] == target {
			return pivot
		} else if target < nums[pivot] {
			right = pivot
		} else {
			left = pivot + 1
		}
	}
	return -1
}