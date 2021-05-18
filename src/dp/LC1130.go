package main

func mctFromLeafValues(arr []int) int {

	n := len(arr)
	max := make([][]int, n)
	for k,_ := range max {
		max[k] = make([]int, n)
	}
	for i := n-1; i >= 0; i-- {
		max[i][i] = arr[i]
		for j:= i+1; j < n; j++ {
			max[i][j] = maxNum(max[i][j-1], arr[j])
		}
	}
	//fmt.Println(max)

	dp := make([][]int, n)
	for k,_ := range dp {
		dp[k] = make([]int, n)
	}
	for i:= n-1; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			dp[i][j] = 0x3f3f3f3f
			for k:=i; k<j; k++ {
				dp[i][j] = minNum(dp[i][j], dp[i][k]+dp[k+1][j]+max[i][k]*max[k+1][j])
			}
		}
	}
	
	return dp[0][n-1]
}

func maxNum(i int, j int) int {
	if i>j {
		return i
	}
	return j
}

func minNum(i int, j int) int {
	if i>j {
		return j
	}
	return i
}

func main() {
	arr := []int{6,2,4}
	mctFromLeafValues(arr)
}