package array

import (
	ds "dsa/datastructures"
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

	openParenthesisStack := ds.NewStack[struct{}]()
	asterisksStack := ds.NewStack[struct{}]()

	for _, char := range input {
		switch string(char) {
		case ")":
			if _, err := openParenthesisStack.Pop(); err == nil {
				break
			}

			if _, err := asterisksStack.Pop(); err == nil {
				break
			}

			return false

		case "(":
			openParenthesisStack.Push(struct{}{})
		case "*":
			asterisksStack.Push(struct{}{})
		}
	}

	canCloseOpenParenthesis := asterisksStack.Len()-openParenthesisStack.Len() >= 0

	return openParenthesisStack.IsEmpty() || canCloseOpenParenthesis
}
