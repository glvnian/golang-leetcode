package day_08

func pathSum(root *TreeNode, targetSum int) int {
	var sumMap = map[int]int{}
	var path = 0
	sumMap[0] = 1
	return pathSumDfs(root, targetSum, sumMap, path)
}

func pathSumDfs(root *TreeNode, targetSum int, sumMap map[int]int, path int) int {
	var count int
	if root == nil {
		return 0
	}
	path += root.Val

	if _, ok := sumMap[path-targetSum]; ok {
		count = sumMap[path-targetSum]
	}

	if _, ok := sumMap[path]; !ok {
		sumMap[path] = 0
	}
	sumMap[path] = sumMap[path] + 1

	count += pathSumDfs(root.Left, targetSum, sumMap, path)
	count += pathSumDfs(root.Right, targetSum, sumMap, path)
	sumMap[path]--

	return count
}

// 暴力
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSumV1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSumV1(root.Left, targetSum)
	res += pathSumV1(root.Right, targetSum)
	return res
}

func pathSumV2(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var pdfs func(*TreeNode, int64)
	pdfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		pdfs(node.Left, curr)
		pdfs(node.Right, curr)
		preSum[curr]--
		return
	}
	pdfs(root, 0)
	return
}
