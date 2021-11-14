package hash

import "strings"

/*
实现一个 MapSum 类，支持两个方法，insert 和 sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
 

示例：

输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)


 */

type MapSum struct {
	Key2Value map[string]int
}


func Constructor() MapSum {
	m := make(map[string]int)
	mapSum := MapSum{
		Key2Value: m,
	}
	return mapSum
}


func (this *MapSum) Insert(key string, val int)  {
	this.Key2Value[key] = val
}


func (this *MapSum) Sum(prefix string) int {
	ret := 0
	for key, val := range this.Key2Value {
		if strings.HasPrefix(key, prefix) {
			ret += val
		}
	}
	return ret
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
