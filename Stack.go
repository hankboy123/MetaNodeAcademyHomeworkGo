package metanodeacademyhomeworkgo

// 使用切片作为栈
type Stack []rune

// 压栈（Push）
func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

// 出栈（Pop）
func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	// 获取最后一个元素
	index := len(*s) - 1
	element := (*s)[index]
	// 删除最后一个元素
	*s = (*s)[:index]
	return element, true
}

// 查看栈顶元素（Peek/Top）
func (s *Stack) Peek() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 获取栈大小
func (s *Stack) Size() int {
	return len(*s)
}

// 清空栈
func (s *Stack) Clear() {
	*s = (*s)[:0]
	// 或者 *s = nil  // 完全释放内存
}
