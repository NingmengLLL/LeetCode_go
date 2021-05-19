package tree

func levelOrder(root *TreeNode) [][]int {

	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		var curRes []int
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			curNode := q[j]
			curRes = append(curRes, curNode.Val)
			if curNode.Left != nil {
				p = append(p, curNode.Left)
			}
			if curNode.Right != nil {
				p = append(p, curNode.Right)
			}
		}
		res = append(res, curRes)
		q = p
	}
	return res
}
