package day_10

type BitTire struct {
	children [2]*BitTire
}

func findMaximumXOR(nums []int) int {
	root := buildNumTire(nums)
	return getMaxXOR(root, nums)
}

func buildNumTire(nums []int) *BitTire {
	var root = &BitTire{}

	for _, num := range nums {
		var node = root
		for i := 31; i >= 0; i-- {
			var bit = (num >> i) & 1
			if node.children[bit] == nil {
				node.children[bit] = &BitTire{}
			}
			node = node.children[bit]
		}
	}
	return root
}

func getMaxXOR(root *BitTire, nums []int) int {
	var max int

	for _, num := range nums {

		var node = root
		var xor = 0
		for i := 31; i >= 0; i-- {
			var bit = (num >> i) & 1
			if node.children[1-bit] != nil {
				xor = (xor << 1) + 1
				node = node.children[1-bit]
			} else {
				xor = (xor << 1)
				node = node.children[bit]
			}
		}
		max = Max(max, xor)
	}

	return max

}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
