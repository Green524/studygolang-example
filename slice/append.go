package slice

import (
	"unicode"
)

func appendInt(x []int, e ...int) []int {
	var z []int
	zlen := len(x) + len(e)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < len(x)*2 {
			zcap = len(x) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	//z[len(x)] = e
	copy(z[len(x):], e)
	return z
}

func remove(x []int, i int) []int {
	copy(x[i:], x[i+1:])
	return x[:len(x)-1]
}

//练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse1(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func rotate(s []int, n int) (rs []int) {
	t1 := s[:n]
	t2 := s[n:]
	for i := 0; i < len(s); i++ {
		if i >= len(t2) {
			rs = append(rs, t1[i-len(t2)])
		} else {
			rs = append(rs, t2[i])
		}
	}
	return
}

func rotateRight(s []int, n int) (rs []int) {
	t1 := s[:len(s)-n]
	t2 := s[len(s)-n:]
	for i := 0; i < len(s); i++ {
		if i >= len(t2) {
			rs = append(rs, t1[i-len(t2)])
		} else {
			rs = append(rs, t2[i])
		}
	}
	return
}

//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func repeat(str []string) []string {
	for i, j := 0, 1; j < len(str); i, j = i+1, j+1 {
		if str[i] == str[j] {
			copy(str[j:], str[j+1:])
			str = str[:len(str)-1]
			i, j = i-1, j-1
		}
	}
	return str
}

//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func space(b []byte) []byte {
	for i, j := 0, 1; j < len(b); i, j = i+1, j+1 {
		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[j])) {
			copy(b[j:], b[j+1:])
			b = b[:len(b)-1]
			i, j = i-1, j-1
		}
	}
	return b
}

//练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
