/*
 * Q.Given a string s containing only three types of characters:
 *   '(', ')' and '*', return true if s is valid.
 *   The following rules define a valid string:
 * - Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 * - Any right parenthesis ')' must have a corresponding left parenthesis '('.
 * - Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 * - '*' could be treated as a single right parenthesis ')'
 *   or a single left parenthesis '('
 *   or an empty string "".
 */

package array

import (
	ds "dsa/datastructures"
	"fmt"
)

func getStarReplacedString(closingBracketCount, openingBracketCount, countStar int, front, raw string) (int, int, string) {
	if len(raw) == 0 {
		return closingBracketCount, openingBracketCount, raw
	}
	head := string(raw[0])
	if head == "(" {
		openingBracketCount++
	} else if head == ")" {
		closingBracketCount++
	} else {
		countStar++
	}
	// fmt.Printf("%2v %2v %2v %v %v\n", closingBracketCount, openingBracketCount, countStar, head, b)
	closingBracketCount, openingBracketCount, processed := getStarReplacedString(closingBracketCount, openingBracketCount, countStar, head, raw[1:])
	if closingBracketCount > openingBracketCount && head == "*" && countStar > 0 {
		processed = "(" + processed
	} else if closingBracketCount < openingBracketCount && head == "*" && countStar > 0 {
		processed = ")" + processed
	} else if closingBracketCount == openingBracketCount && head == "*" && countStar > 0 {
		processed = "" + processed
	} else {
		processed = head + processed
	}
	// fmt.Printf("%2v %2v %2v %v %v\n", closingBracketCount, openingBracketCount, countStar, head, b)
	return closingBracketCount, openingBracketCount, processed
}

func RunCheckMatchingBrackets() {
	input := []string{
		"()()()()()()()()()()(*))", // valid
		"(*))()()()()()()()()()()", // valid
		"(((*)))",                  // valid
		"()())*",                   // invalid
		"()*)",                     // valid
		"(((*)",                    // invalid
	}
	for index := 0; index < len(input); index++ {
		_, _, final := getStarReplacedString(0, 0, 0, "", input[index])
		fmt.Println("Processed String:", final)
		if len(final)%2 != 0 {
			fmt.Println("Ans: Is odd and invalid")
		} else {
			stack := ds.InitStack("")
			for i := 0; i < len(final); i++ {
				val := string(final[i])
				// fmt.Println("stack: ", i, stack.Top())
				if stack.Top() == "(" && val == ")" {
					stack.Pop()
				} else {
					stack.Push(val)
				}
			}
			// fmt.Println(stack.GetStack())
			if stack.IsEmpty() {
				fmt.Println("Ans: Is even and valid")
			} else {
				fmt.Println("Ans: Is even and invalid")
			}
		}
		fmt.Println()
	}
}
