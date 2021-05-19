package tree

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

	var res bool
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {

		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if node.Val == x {
			if left > n/2 || right > n/2 || 1 + left + right <= n/2 {
				res = true
			}
		}
		return 1 + left +right
	}
	dfs(root)
	return res
}
