package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	org, err := clipboard.ReadAll()
	t := ChainableString(org)

	if err != nil {
		fmt.Println("panic on read clipboard: ", err)
		return
	}

	if err := clipboard.WriteAll(string(convert(t))); err != nil {
		fmt.Println("panic on write clipboard: ", err)
	}
}

// NOTE: 必要最低限。ブラケット、例外は各自
func convert(t ChainableString) ChainableString {
	return t.
		Replace("String", "string").
		Replace("boolean", "bool").
		Replace("System.out.", "Console.").
		Replace("println", "WriteLine").
		Replace("```java", "```cs").
		Replace(".length", ".Length").
		Replace("ArrayIndexOutOfBound", "IndexOutOfRange")
}
