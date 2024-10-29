package main

import (
	"flag"
	"fmt"
	"strings"

	workdirs "github.com/Chaitanyabsprip/workdirs/pkg"
)

func main() {
	short := flag.Bool("s", false, "short")
	flag.Parse()
	if *short {
		fmt.Println(
			strings.Join(workdirs.Shorten(workdirs.Worktrees()), "\n"),
		)
		return
	}
	fmt.Println(strings.Join(workdirs.Worktrees(), "\n"))
}
