package main

import (
	"fmt"
	"strings"

	workdirs "github.com/Chaitanyabsprip/workdirs/pkg"
)

func main() {
	fmt.Println(strings.Join(workdirs.Worktrees(), "\n"))
}
