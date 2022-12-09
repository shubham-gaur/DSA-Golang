package array

import (
	ds "github.com/shubham-gaur/dsa-golang/datastructures"
)

/*
 * ValidateParenthesisStack take a string containing only three types of characters:
 * '(', ')' and '*', return true if s is valid.
 *
 * The following rules define a valid string:
 *  - Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 *  - Any right parenthesis ')' must have a corresponding left parenthesis '('.
 *  - Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 *  - '*' could be treated as ')' or '(' or "".
 */
func ValidateParenthesisStack(input string) bool {
	if len(input) == 0 {
		return true
	}

	// keep count of the open parenthesis and the asterisks
	// TODO: instead of a stack, we can keep count of both variables as int
	openParenthesisStack := ds.NewStack[struct{}]()
	asterisksStack := ds.NewStack[struct{}]()

	for _, char := range input {
		switch string(char) {
		case ")":
			// check if we have an open parenthesis to close
			if _, err := openParenthesisStack.Pop(); err == nil {
				break
			}

			// check if we can use an asterisks to close with this parenthesis
			if _, err := asterisksStack.Pop(); err == nil {
				break
			}

			// we do not have enough parenthesis or asterisks to close
			return false

		case "(":
			// increase the open parenthesis count
			openParenthesisStack.Push(struct{}{})
		case "*":
			// increase the asterisks count
			asterisksStack.Push(struct{}{})
		}
	}

	// check if we have enough asterisks to close the open parenthesis left
	canCloseOpenParenthesis := asterisksStack.Len()-openParenthesisStack.Len() >= 0

	// check if we have some open parenthesis left
	return openParenthesisStack.IsEmpty() || canCloseOpenParenthesis
}
