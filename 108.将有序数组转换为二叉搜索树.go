package cmd

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) - 1) / 2
	tn := &TreeNode{Val: nums[mid]}
	tn.Left = sortedArrayToBST(nums[:mid])
	tn.Right = sortedArrayToBST(nums[mid+1:])
	return tn
}

// @lc code=end
