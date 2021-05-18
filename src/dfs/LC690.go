package dfs

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	res := 0

	var dfs func(i int)
	dfs = func(i int) {
		for _, e := range employees{
			if e.Id == i {
				res += e.Importance
				for _, subId := range e.Subordinates {
					dfs(subId)
				}
			}
		}

	}
	dfs(id)
	return res
}
