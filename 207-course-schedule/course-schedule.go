func canFinish(numCourses int, prerequisites [][]int) bool {
    postrequisitesMap := make([]map[int]bool, numCourses)
    indegree := make([]int, numCourses)
    
    for i := 0; i < numCourses; i++ {
        postrequisitesMap[i] = make(map[int]bool)
    }

    for _, courses := range prerequisites {
        course, prereqCourse := courses[0], courses[1]

        postrequisitesMap[prereqCourse][course] = true
        indegree[course] += 1
    }

    curCourses := []int{}

    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            curCourses = append(curCourses, i)
        }
    }

    for len(curCourses) != 0 {
        newCourses := []int{}

        for _, course := range curCourses {
            for postCourse, _ := range postrequisitesMap[course] {
                indegree[postCourse] -= 1
                
                if indegree[postCourse] == 0 {
                    newCourses = append(newCourses, postCourse)
                }
            }
        }

        curCourses = newCourses
    }

    for _, degree := range indegree {
        if degree > 0 {
            return false
        }
    }

    return true
}