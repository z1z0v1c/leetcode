/**
 * 22 - Medium
 *
 * Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
 *
 * Example 1:
 *
 * 	Input: n = 3
 * 	Output: ["((()))","(()())","(())()","()(())","()()()"]
 *
 * Example 2:
 *
 * 	Input: n = 1
 * 	Output: ["()"]
 *
 * Constraints:
 *
 * 	- 1 <= n <= 8
 */
package generateparentheses

func generateParenthesis(n int) []string {
	ans := []string{}

	var generate func(int, int, string)
	generate = func(left, right int, current string) {
		if len(current) == 2*n {
			ans = append(ans, current)
		}
	
		if left > 0 {
			generate(left-1, right, current + "(")
		}
	
		if right > left {
			generate(left, right-1, current + ")")
		}
	}

	generate(n, n, "")
	
	return ans
}
