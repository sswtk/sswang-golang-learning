package main

import "fmt"

/*
fg: 寻找最长不含有重复字符的字串（leetcode）
	abcabcbb -> abc
	bbbb -> b
	pwwkew -> wk
*/


/*
thinking:
	lastOccurred[x]不存在，或者< start -> 无需操作
	lastOccurred[x] >= start -> 更新start
	更新lastOccurred[x],更新maxLength
 */

func lengthOfNoneRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start, maxLength := 0, 0
	for i, ch := range [] byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >=  start{
			start = lastI + 1
		}
		if i-start+1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main(){
	fmt.Println(lengthOfNoneRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNoneRepeatingSubStr("bbb"))
	fmt.Println(lengthOfNoneRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNoneRepeatingSubStr(""))
	fmt.Println(lengthOfNoneRepeatingSubStr("b"))
	fmt.Println(lengthOfNoneRepeatingSubStr("pfewd"))
}

