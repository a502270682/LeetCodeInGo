package bit

/*
有两种特殊字符：

第一种字符可以用一个比特0来表示
第二种字符可以用两个比特(10或11)来表示、
给定一个以 0 结尾的二进制数组bits，如果最后一个字符必须是一位字符，则返回 true 。



示例1:

输入: bits = [1, 0, 0]
输出: true
解释: 唯一的编码方式是一个两比特字符和一个一比特字符。
所以最后一个字符是一比特字符。

*/

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return bits[0] == 0
	}
	for i := 0; i < len(bits); {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
		if i < len(bits) && i == len(bits)-1 && bits[i] == 0 {
			return true
		}
	}
	return false
}
