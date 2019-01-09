## 题目描述
``
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
``
## 思路
1. 首先比较 nums1 和 nums2 调换下位置，后面遍历小的数组会快一些。
2. 遍历 num1，通过 map 存储 num1 的元素与频率。
3. 遍历 num2，在 map 中查找是否有相同的元素（该元素的存储频率大于0），如果有，用 map 存储，同时该元素的频率减一.
4. 返回数组就是它们的交集。

* 时间复杂度: O(nlogn)
* 空间复杂度: O(n)

```go
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 =  nums2, nums1
	}

	uniqueMap := map[int]int{}
	for _, i := range nums1 {
		if _, ok := uniqueMap[i]; ok {
			uniqueMap[i] += 1
		} else {
			uniqueMap[i] = 1
		}
	}

	var response []int
	for _, j := range nums2 {
		if num, ok := uniqueMap[j]; ok && num > 0 {
			uniqueMap [j] = num - 1
			response = append(response, j)
		}
	}
	return response
}
```