package main

import (
	metanodeacademyhomeworkgo "MetaNodeAcademyHomeworkGo"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, 世界")

	isRight := isValid("{]")
	fmt.Println(isRight)
	singItemIn := []int{1, 2, 3}

	singleItem, err := findSingeItem(singItemIn[:])

	if err == nil {
		fmt.Println(singleItem)
	}
}

func findSingeItem(items []int) (int, error) {
	if len(items) == 0 {
		fmt.Println("切片为空")
		return 0, fmt.Errorf("元素为空!")
	}
	counts := make(map[int]int)
	for _, v := range items {
		counts[v]++
	}
	if len(counts) == 0 {
		return 0, fmt.Errorf("元素为空！")
	}

	for key, value := range counts {
		if value == 1 {
			return key, nil
		}
	}

	return 0, fmt.Errorf("没有找到只有一个值的元素！")
}

func isPalindrome(item int) bool {
	if item <= 0 || (item%10 == 0 && item != 0) {
		return false
	}
	reverse := reverse(item)
	return item == reverse
}

func reverse(item int) int {

	reverse := 0
	for item > 0 {
		reverse = reverse*10 + reverse%10
		item /= 10
	}
	return reverse

}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func isValid(s string) bool {

	if len(s) == 0 {
		return false
	}
	pairMap := make(map[rune]rune)
	pairMap[')'] = '('
	pairMap['}'] = '{'
	pairMap[']'] = '['
	pairMap['('] = '0'
	pairMap['{'] = '0'
	pairMap['['] = '0'
	var stack metanodeacademyhomeworkgo.Stack

	for _, ch := range s {
		if value, ok := pairMap[ch]; ok {
			if value == '0' {
				stack.Push(ch)
			} else {
				popChar, exist1 := stack.Pop()
				if !exist1 {
					return false
				}
				if popChar == value {
					continue
				} else {
					return false
				}
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}


findLongestPre(strs []string) string{
	if len(strs) ==0 {
		return ""
	}

	prefix := strs[0]

	for i:=1;i< len(strs);i++{
		j:=0
		for j<len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j]{
			j++
		}

		prefix = prefix[:j]
		if prefix == ""{
			return ""
		}
	}
}


func plusOne(digits []int) []int {
    n := len(digits)
    
    // 从最后一位开始向前遍历
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        // 当前位是9，加1后变为0，产生进位
        digits[i] = 0
    }
    
    // 如果所有位都是9（如999->1000），需要在最前面加1
    result := make([]int, n+1)
    result[0] = 1
    // 后面的位都是0
    return result
}

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。

判题标准:

系统会用下面的代码来测试你的题解:
int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案
int k = removeDuplicates(nums); // 调用
assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。

*/
func removeDuplicates(nums []int) int {
     if len(nums) ==0 {
		return 0
	}

	i :=0
	for j :=0;j<len(nums);j++{
		
		if nums[j]!=nums[i]{
			nums[i+1]=nums[j]
            i++

		}
	}
    return i+1
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge( [][2]int intervals) [][2]int{

}


/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

 

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
 

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/
func twoSum(nums []int, target int) []int {
	 // 创建一个哈希表来存储数字和对应的索引
    numMap := make(map[int]int)
    
    for i, num := range nums {
        // 计算目标值与当前数字的差值
        complement := target - num
        
        // 检查差值是否已经在哈希表中
        if index, found := numMap[complement]; found {
            // 找到答案，返回两个索引
            return []int{index, i}
        }
        
        // 将当前数字和索引添加到哈希表中
        numMap[num] = i
    }
    
    // 根据题目假设，总会有一个答案，所以这里不会执行到
    return nil
    
    
}