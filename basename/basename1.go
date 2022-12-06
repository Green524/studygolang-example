package basename

import (
	"bytes"
	"strings"
)

func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//111123456.111
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
//练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
func commaBuf(s string) string {
	var buf bytes.Buffer
	//111,222,333
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		buf.WriteString(s[0:1])
		s = s[1:]
	}
	for i, e := range s {
		if e == '.' {
			buf.WriteString(s[i:])
			break
		}
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(byte(e))
	}
	return buf.String()
}

//练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func dis(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}
	//这个循环判断拥有相同字符
	for _, e := range s1 {
		if !strings.ContainsRune(s2, e) {
			return false
		}
	}
	//到达这里都是有着相同的字符，然后判断顺序
	for i, e := range s1 {
		//走完循环都没有进入if，说明顺序相同
		if e != rune(s2[i]) {
			return true
		}
	}

	return false
}
