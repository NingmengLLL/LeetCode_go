package binary

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var ans [][]int
	for i:=1;i<=1000;i++ {
		if customFunction(i,1) > z {
			break
		}
		for l,r := 1,1000;l<=r;{
			mid := l + (r-l)>>1
			val := customFunction(i,mid)
			if val > z {
				r = mid-1
			}else if val < z {
				l = mid+1
			}else{
				ans = append(ans, []int{i,mid})
				break
			}
		}
	}
	return ans
}
