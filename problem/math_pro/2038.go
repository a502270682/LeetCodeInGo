package math_pro

/*
总共有 n个颜色片段排成一列，每个颜色片段要么是'A'要么是'B'。给你一个长度为n的字符串colors，其中colors[i]表示第i个颜色片段的颜色。

Alice 和 Bob 在玩一个游戏，他们 轮流从这个字符串中删除颜色。Alice 先手。

如果一个颜色片段为 'A'且 相邻两个颜色都是颜色 'A'，那么 Alice 可以删除该颜色片段。Alice不可以删除任何颜色'B'片段。
如果一个颜色片段为 'B'且 相邻两个颜色都是颜色 'B'，那么 Bob 可以删除该颜色片段。Bob 不可以删除任何颜色 'A'片段。
Alice 和 Bob 不能从字符串两端删除颜色片段。
如果其中一人无法继续操作，则该玩家 输掉游戏且另一玩家 获胜。
假设 Alice 和 Bob 都采用最优策略，如果 Alice 获胜，请返回true，否则 Bob 获胜，返回false。



示例 1：

输入：colors = "AAABABB"
输出：true
解释：
AAABABB -> AABABB
Alice 先操作。
她删除从左数第二个 'A' ，这也是唯一一个相邻颜色片段都是 'A' 的 'A' 。

现在轮到 Bob 操作。
Bob 无法执行任何操作，因为没有相邻位置都是 'B' 的颜色片段 'B' 。
因此，Alice 获胜，返回 true 。
*/

/*
题解
alice: 找到3个以上连续的A，减去A-2，这个个数就是alice可以操作的步数
同理，bob也是连续的B，步数=连续的B数目-2
*/

func winnerOfGame(colors string) bool {
	freq := [2]int{}
	cur, cnt := 'C', 0
	for _, c := range colors {
		if c != cur {
			cur, cnt = c, 1
		} else {
			cnt++
			if cnt >= 3 {
				freq[cur-'A']++
			}
		}
	}
	return freq[0] > freq[1]
}
