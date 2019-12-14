package leetcode

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
分两步
1、所有的非0往前移，
2、后面补0
 */
func MoveZeroes(nums []int)  {
	var i = 0

	for _,v := range nums{
		if v != 0 {
			nums[i] = v
			i ++
		}
	}

	for j := i;j<len(nums);j ++{
		nums[j] = 0
	}

}

func MoveZeroesSolutionTwo(nums []int)  {
	var j = 0

	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			if i != j{
				nums[j],nums[i] = nums[i],nums[j]
			}
			j++
		}
	}
}
