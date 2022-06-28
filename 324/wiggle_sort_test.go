package _324

import (
	"sort"
	"testing"
)

// 解题思路
// 先排序，再通过双指针法实现数组重排
func TestWiggleSort(t *testing.T) {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	t.Log("nums:", nums)
}

func wiggleSort(nums []int) {
	// 获取数组长度
	n := len(nums)
	// 生成一个新数组
	arr := append([]int{}, nums...)
	// 将新数组进行排序，使用sort包中的sort.Ints()
	sort.Ints(arr)
	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}
