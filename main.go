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
