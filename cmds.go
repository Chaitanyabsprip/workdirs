package workdirs

import (
	"fmt"
	"os"
	"strings"

	bonzai "github.com/rwxrob/bonzai/pkg"
	"github.com/rwxrob/bonzai/pkg/core/comp"
)

var Cmd = &bonzai.Cmd{
	Name: `work`,
	Cmds: []*bonzai.Cmd{dirsCmd, treeCmd},
	Comp: comp.Cmds,
	Call: func(x *bonzai.Cmd, args ...string) error {
		return nil
	},
}

var dirsCmd = &bonzai.Cmd{
	Name:  `dirs`,
	Alias: `d`,
	Call: func(x *bonzai.Cmd, args ...string) error {
		short := len(os.Getenv(`SHORT`)) > 0 ||
			len(os.Getenv(`W_SHORT`)) > 0
		var out string
		dirs := make([]string, 0)
		dirs = append(dirs, Workdirs()...)
		dirs = append(dirs, Worktrees()...)
		if !short {
			out = strings.Join(dirs, "\n")
		} else {
			out = strings.Join(Shorten(dirs), "\n")
		}
		fmt.Println(out)

		return nil
	},
}

var treeCmd = &bonzai.Cmd{
	Name:  `trees`,
	Alias: `t`,
	Call: func(x *bonzai.Cmd, args ...string) error {
		short := len(os.Getenv(`SHORT`)) > 0 ||
			len(os.Getenv(`W_SHORT`)) > 0
		var out string
		if !short {
			out = strings.Join(Worktrees(), "\n")
		} else {
			out = strings.Join(
				Shorten(Worktrees()),
				"\n",
			)
		}
		fmt.Println(out)
		return nil
	},
}
