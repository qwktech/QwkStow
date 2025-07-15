package main

import (
	"fmt"
	"regexp"
	"strings"
)

// "regexp"
// "strings"

func RegexpTest(target string) string {
	var result string
	parts := strings.Split(target, "/")
	fileName := parts[len(parts)-1]
	updated := parts[:len(parts)-1]

	if strings.HasPrefix(fileName, "dot-") {
	       strings.Replace(fileName, "dot-", ".", 1))
	}

	return strings.Join(parts, "/")
}

func main() {
	target := "/home/Hayden/dotfiles/dot-vimrc"
	parts := strings.Split(target, "/")

	fmt.Println("len: ", len(parts))
	fmt.Println("last: ", parts[(len(parts)-1)])

	fmt.Println("replacement: ", 
}
