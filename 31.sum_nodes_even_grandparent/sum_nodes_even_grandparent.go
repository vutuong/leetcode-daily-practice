package sumEvenGrandparent

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/

1315. Sum of Nodes with Even-Valued Grandparent

Given a binary tree, return the sum of values of nodes with even-valued grandparent.
A grandparent of a node is the parent of its parent, if it exists.
If there are no nodes with an even-valued grandparent, return 0.
*/

// Define the result type
type result struct {
	Val int
}

func sumEvenGrandparent(root *TreeNode) int {
	cur := root
	r := new(result)
	r.addGrandChild(cur)
	return r.Val
}

// Function walk through node in a tree and update result.Val as sum of every nodes with an even-valued grandparent
func (r *result) addGrandChild(cur *TreeNode) {
	if cur != nil {
		if cur.Val%2 == 0 {
			if cur.Left != nil {
				if cur.Left.Left != nil {
					r.Val += cur.Left.Left.Val
				}
				if cur.Left.Right != nil {
					r.Val += cur.Left.Right.Val
				}
				r.addGrandChild(cur.Left)
			}
			if cur.Right != nil {
				if cur.Right.Left != nil {
					r.Val += cur.Right.Left.Val
				}
				if cur.Right.Right != nil {
					r.Val += cur.Right.Right.Val
				}
				r.addGrandChild(cur.Right)
			}
		} else {
			r.addGrandChild(cur.Left)
			r.addGrandChild(cur.Right)
		}
	}
}
