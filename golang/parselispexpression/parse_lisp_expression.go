/**
 * 736 - Hard
 *
 * You are given a string expression representing a Lisp-like expression to return the integer value of.
 *
 * The syntax for these expressions is given as follows.
 *
 * 	- An expression is either an integer, let expression, add expression, mult expression, or an assigned variable.
 *    Expressions always evaluate to a single integer. (An integer could be positive or negative.)
 *
 * 	- A let expression takes the form "(let v1 e1 v2 e2 ... vn en expr)", where let is always the string "let",
 *    then there are one or more pairs of alternating variables and expressions, meaning that the first variable v1
 * 	  is assigned the value of the expression e1, the second variable v2 is assigned the value of the expression e2,
 *	  and so on sequentially; and then the value of this let expression is the value of the expression expr.
 *
 * 	- An add expression takes the form "(add e1 e2)" where add is always the string "add", there are always two expressions
 *    e1, e2 and the result is the addition of the evaluation of e1 and the evaluation of e2.
 *
 * 	- A mult expression takes the form "(mult e1 e2)" where mult is always the string "mult", there are always two expressions
 *    e1, e2 and the result is the multiplication of the evaluation of e1 and the evaluation of e2.
 *
 * 	- For this question, we will use a smaller subset of variable names. A variable starts with a lowercase letter,
 *    then zero or more lowercase letters or digits. Additionally, for your convenience, the names "add", "let", and "mult"
 *    are protected and will never be used as variable names.
 *
 * 	- Finally, there is the concept of scope. When an expression of a variable name is evaluated, within the context of that
 *    evaluation, the innermost scope (in terms of parentheses) is checked first for the value of that variable, and then outer
 *    scopes are checked sequentially. It is guaranteed that every expression is legal. Please see the examples for more details
 *    on the scope.
 *
 * Example 1:
 *
 *	Input: expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
 *	Output: 14
 *	Explanation:
 *
 *		In the expression (add x y), when checking for the value of the variable x,
 * 		we check from the innermost scope to the outermost in the context of the variable we are trying to evaluate.
 * 		Since x = 3 is found first, the value of x is 3.
 *
 * Example 2:
 *
 * 	Input: expression = "(let x 3 x 2 x)"
 * 	Output: 2
 * 	Explanation: Assignment in let statements is processed sequentially.
 *
 * Example 3:
 *
 * 	Input: expression = "(let x 1 y 2 x (add x y) (add x y))"
 * 	Output: 5
 * 	Explanation:
 *
 *		The first (add x y) evaluates as 3, and is assigned to x.
 * 		The second (add x y) evaluates as 3+2 = 5.
 *
 * Constraints:
 *
 * 	- 1 <= expression.length <= 2000
 * 	- There are no leading or trailing spaces in expression.
 * 	- All tokens are separated by a single space in expression.
 * 	- The answer and all intermediate calculations of that answer are guaranteed to fit in a 32-bit integer.
 * 	- The expression is guaranteed to be legal and evaluate to an integer.
 */
package parselispexpression

import (
	"strconv"
	"strings"
)

func evaluate(expression string) int {
	return parse(expression, map[string]int{})
}

// Parse expression recursively
func parse(expression string, context map[string]int) int {
	// Return converted value if expression is a number
	if isDigit(expression[0]) || expression[0] == '-' {
		result, _ := strconv.Atoi(expression)
		return result
	}

	// Return value from the context if expression is a variable
	if isLetter(expression[0]) {
		return context[expression]
	}

	// Trim parentheses
	expression = expression[1 : len(expression)-1]

	// Cut the first word as operation
	operation, expression, _ := strings.Cut(expression, " ")

	switch operation {

	case "let":
		tokens := tokenize(expression)

		for i := 0; i < len(tokens)-1; i = i + 2 {
			if tokens[i+1][0] == '(' {
				context[tokens[i]] = parse(tokens[i+1], copy(context))
			} else {
				val, err := strconv.Atoi(tokens[i+1])
				if err != nil {
					context[tokens[i]] = parse(tokens[i+1], copy(context))
				} else {
					context[tokens[i]] = val
				}
			}
		}

		return parse(tokens[len(tokens)-1], copy(context))

	case "add":
		sum := 0
		tokens := tokenize(expression)

		for _, token := range tokens {
			if isDigit(token[0]) || token[0] == '-' {
				val, _ := strconv.Atoi(token)
				sum += val
			} else if token[0] == '(' {
				sum += parse(token, copy(context))
			} else {
				sum += context[token]
			}
		}

		return sum

	case "mult":
		prod := 1
		tokens := tokenize(expression)

		for _, token := range tokens {
			if isDigit(token[0]) || token[0] == '-' {
				val, _ := strconv.Atoi(token)
				prod *= val
			} else if token[0] == '(' {
				prod *= parse(token, copy(context))
			} else {
				prod *= context[token]
			}
		}

		return prod
	}

	return 0
}

// Parse tokens from an expression
func tokenize(expression string) []string {
	// Count open parentheses to find subexpressions 
	parentheses := 0
	// Keep track of the end of previous token
	prev := 0

	tokens := make([]string, 0)

	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			parentheses++
		} else if expression[i] == ')' {
			parentheses--
			
			if parentheses == 0 {
				token := ""

				if i == len(expression)-1 {
					token = expression[prev:]
				} else if expression[i+1] != ' ' {
					token = expression[prev:i]
				}

				if token != "" {
					tokens = append(tokens, token)
					prev = i + 1
				}
			}
		} else if (expression[i] == ' ' && parentheses == 0) || i == len(expression)-1 {
			token := ""

			if i == len(expression)-1 {
				token = expression[prev:]
			} else {
				token = expression[prev:i]
			}

			tokens = append(tokens, token)
			prev = i + 1
		}
	}

	return tokens
}

// Check if character is a digit
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

// Check if character is a letter
// It is guaranteed to be a lowercase 
func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z'
}

// Copy map
func copy(origin map[string]int) map[string]int {
	copy := make(map[string]int, len(origin))

	for k, v := range origin {
		copy[k] = v
	}

	return copy
}
