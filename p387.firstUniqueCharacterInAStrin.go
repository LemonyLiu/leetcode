package leetcode

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func FirstUniqChar(s string) int {
	strMap := make(map[rune]int)

	for index,v := range s{
		if _,ok := strMap[v];ok{
			strMap[v] = -1
		} else {
			strMap[v] = index
		}
	}
	// map 无序，这里会出错
	for _,v := range strMap{
		if v != -1{
			return v
		}
	}
	return -1
}

func FirstUniqChar2(s string) int {
	strMap := make(map[rune]int)
	var strList []rune
	for index,v := range s{
		if _,ok := strMap[v];ok{
			strMap[v] = -1
		} else {
			strList = append(strList, v)
			strMap[v] = index
		}
	}
	for _,v := range strList{
		if index,ok :=strMap[v];ok{
			if index != -1 {
				return index
			}
		}
	}
	return -1
}