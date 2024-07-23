package bfs2

// https://www.lintcode.com/problem/615/

/**
 * @param numCourses: a total of n courses
 * @param prerequisites: a list of prerequisite pairs
 * @return: true if can finish all courses or false
 */
func CanFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make(map[int]int)
	dependencies := make(map[int][]int)
	for _, pair := range prerequisites {
		indegree[pair[0]]++
		if _, ok := indegree[pair[1]]; !ok {
			indegree[pair[1]] = 0
		}
		dependencies[pair[1]] = append(dependencies[pair[1]], pair[0])
	}

	queue := make([]int, 0)
	for course, cnt := range indegree {
		if cnt == 0 {
			queue = append(queue, course)
		}
	}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		for _, v := range dependencies[head] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, cnt := range indegree {
		if cnt > 0 {
			return false
		}
	}
	return true
}
