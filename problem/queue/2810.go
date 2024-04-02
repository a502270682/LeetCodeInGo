package queue

func finalString(s string) string {
	q := []rune{}
	head := false
	for _, i := range s {
		if i == 'i' {
			head = !head
		} else {
			if head {
				// 加在队列头
				q = append([]rune{i}, q...)
			} else {
				// 加在队列尾
				q = append(q, i)
			}
		}
	}
	if head {
		for i, j := 0, len(q)-1; i < j; i, j = i+1, j-1 {
			q[i], q[j] = q[j], q[i]
		}
	}
	return string(q)
}
