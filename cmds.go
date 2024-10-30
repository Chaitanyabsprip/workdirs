package workdirs

import (
	"fmt"
	"os"
	"strings"

	bonzai "github.com/rwxrob/bonzai/pkg"
	"github.com/rwxrob/bonzai/pkg/core/comp"
)

var Cmd = &bonzai.Cmd{
	Name:    `work`,
	Usage:   `[SHORT=1] work dirs|trees`,
	Short:   `List work directories and git worktree repositories. Set SHORT for compact output`,
	Default: helpCmd,
	Cmds:    []*bonzai.Cmd{helpCmd, dirsCmd, treeCmd},
	Comp:    comp.Cmds,
}

var helpCmd = &bonzai.Cmd{
	Name:  `help`,
	Usage: `work help`,
	Short: `Show this help message`,
	Call: func(x *bonzai.Cmd, args ...string) error {
		cmds := []*bonzai.Cmd{dirsCmd, treeCmd}
		switch len(args) {
		case 0:
			for _, cmd := range cmds {
				fmt.Println(cmd.Usage)
				fmt.Println("\t", cmd.Short)
				fmt.Println("\t", cmd.Long)
			}
		case 1:
			var matched bool
			for _, cmd := range cmds {
				if cmd.Name == args[0] {
					matched = true
					fmt.Println(cmd.Usage)
					fmt.Println("\t", cmd.Short)
					fmt.Println("\t", cmd.Long)
				}
			}
			if !matched {
				return fmt.Errorf(
					"unknown command %q",
					args[0],
				)
			}
		default:
			return fmt.Errorf("too many arguments")
		}
		return nil
	},
}

var dirsCmd = &bonzai.Cmd{
	Name:  `dirs`,
	Alias: `d`,
	Usage: `[SHORT=1] dirs`,
	Short: `List work directories. Set SHORT for compact output`,
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
	Usage: `[SHORT=1] trees`,
	Short: `List git worktrees. Set SHORT for compact output`,
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
