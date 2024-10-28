package main

import (
	"fmt"
	"strings"

	workdirs "github.com/Chaitanyabsprip/workdirs/pkg"
)

func main() {
	dirs := make([]string, 0)
	dirs = append(dirs, workdirs.Workdirs()...)
	dirs = append(dirs, workdirs.Worktrees()...)
	fmt.Println(strings.Join(dirs, "\n"))
}
