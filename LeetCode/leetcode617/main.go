package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type person struct {
	name string
}

func main() {
	var r1, r2 *TreeNode
	r1 = &TreeNode{
		1,
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		&TreeNode{
			Val: 2,
		},
	}
	r2 = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tmp := mergeTrees(r1, r2)
	fmt.Println(tmp)
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
