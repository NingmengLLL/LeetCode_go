package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	ans := make([]int, k+1)
	m := make(map[int]map[int]struct{})

	for _,v := range logs {
		if _,ok := m[v[0]]; !ok {
			m[v[0]] = map[int]struct{}{v[1]: {}}
			continue
		}
		m[v[0]][v[1]] = struct{}{}
	}

	for _,val := range m {
		ans[len(val)]++
	}
	return ans[1:]
}