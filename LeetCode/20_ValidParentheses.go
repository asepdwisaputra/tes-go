package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))           // true
	fmt.Println(isValid("()[]{}"))       // true
	fmt.Println(isValid("(]"))           // false
	fmt.Println(isValid("([)]"))         // false
	fmt.Println(isValid("{[]}"))         // true
	fmt.Println(isValid("{{{{{{{{{{{{")) // false
}
