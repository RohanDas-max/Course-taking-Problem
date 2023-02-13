package main

func main() {

}

func coursesToTake(arrs [][]int, n int) []int {
	if len(arrs) <= 0 {
		return []int{0}
	} else {
		res := make([]int, 0, n)
		for _, arr := range arrs {
			for i := (len(arr) - 1); i >= 0; i-- {
				if len(res) > 0 && !contains(res, arr[i]) {
					res = append(res, arr[i])
				} else if len(res) <= 0 {
					res = append(res, arr[i])
				}
			}
		}
		return res
	}
}

func contains(arr []int, b int) bool {
	for _, a := range arr {
		if a == b {
			return true
		}
	}
	return false
}

func findOrder(arrs [][]int, n int) []int {
	graph, incomings := buildMap(n, arrs)

	var starts []int
	for i, incoming := range incomings {
		if incoming == 0 {
			starts = append(starts, i)
		}
	}

	var res []int
	var visitNum int
	for _, start := range starts {
		queue := []int{start}
		res = append(res, start)
		visitNum += 1
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			neighbors, ok := graph[node]
			if ok {
				for _, neighbor := range neighbors {
					incomings[neighbor] -= 1
					if incomings[neighbor] == 0 {
						queue = append(queue, neighbor)
						res = append(res, neighbor)
						visitNum += 1
					}
				}
			}
		}
	}
	if visitNum == n {
		return res
	}
	return []int{}
}

func buildMap(n int, arrs [][]int) (map[int][]int, []int) {
	graph := make(map[int][]int)
	incomings := make([]int, n)
	for _, preq := range arrs {
		if _, ok := graph[preq[1]]; !ok {
			graph[preq[1]] = []int{preq[0]}
		} else {
			graph[preq[1]] = append(graph[preq[1]], preq[0])
		}
		incomings[preq[0]] += 1
	}
	return graph, incomings
}
