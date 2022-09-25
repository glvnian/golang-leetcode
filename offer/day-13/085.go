package day_13

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	helperParenthesis(n, n, "", &result)
	return result
}

func helperParenthesis(left, right int, subnet string, result *[]string) {

	if left == 0 && right == 0 {
		*result = append(*result, subnet)
	}

	if left > 0 {
		helperParenthesis(left-1, right, subnet+"(", result)
	}

	if left < right {
		helperParenthesis(left, right-1, subnet+")", result)
	}
}
