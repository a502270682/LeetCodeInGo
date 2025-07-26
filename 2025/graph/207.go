package graph

type CourseMsg struct {
	preCourseCount int
	nextVals       []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses < 1 || len(prerequisites) < 1 {
		return true
	}
	course2NextCourses := make(map[int]*CourseMsg)
	for _, prerequisite := range prerequisites {
		course2NextCourses[prerequisite[1]] = &CourseMsg{
			nextVals: make([]int, 0),
		}
		course2NextCourses[prerequisite[0]] = &CourseMsg{
			nextVals: make([]int, 0),
		}

	}
	for _, prerequisite := range prerequisites {
		preVal := prerequisite[1]
		nextVal := prerequisite[0]
		course2NextCourses[nextVal].preCourseCount++
		course2NextCourses[preVal].nextVals = append(course2NextCourses[preVal].nextVals, nextVal)
	}
	preZeroCourses := make([]int, 0)
	for key, course := range course2NextCourses {
		if course.preCourseCount == 0 {
			preZeroCourses = append(preZeroCourses, key)
		}
	}
	for len(preZeroCourses) > 0 {
		length := len(preZeroCourses)
		for i := 0; i < length; i++ {
			preVal := preZeroCourses[i]
			for _, nextVal := range course2NextCourses[preVal].nextVals {
				course2NextCourses[nextVal].preCourseCount--
				if course2NextCourses[nextVal].preCourseCount == 0 {
					preZeroCourses = append(preZeroCourses, nextVal)
				}
			}
			delete(course2NextCourses, preVal)
			numCourses--
		}
		preZeroCourses = preZeroCourses[length:]
	}

	return numCourses == 0 || len(course2NextCourses) == 0
}
