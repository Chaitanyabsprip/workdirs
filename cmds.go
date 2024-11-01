package workdirs

import (
	"fmt"
	"os"
	"strings"

	bonzai "github.com/rwxrob/bonzai/pkg"
	"github.com/rwxrob/bonzai/pkg/core/comp"
)

// Cmd provides access to the `work` command suite, which lists work
// directories and Git worktree repositories. Use SHORT=1 for compact
// output.
var Cmd = &bonzai.Cmd{
	Name:  `work`,
	Usage: `work <command>`,
	Short: `List work directories and Git worktrees`,
	Long: `The 'work' command lists local work directories and
Git worktree repositories. It supports compact output when SHORT=1 is
set.`,
	Default: helpCmd,
	Cmds:    []*bonzai.Cmd{helpCmd, dirsCmd, treeCmd},
	Comp:    comp.Cmds,
}

// helpCmd shows detailed help for each subcommand or a summary of all
// commands.
var helpCmd = &bonzai.Cmd{
	Name:   `help`,
	Usage:  `work help [command]`,
	Short:  `show help for a command`,
	Params: `dirs|trees`,
	Comp:   comp.Params,
	Long: `
Displays usage and description for all available commands, or detailed
information about a specified command if provided.

USAGE
  work help [command]

EXAMPLES
  work help       # shows usage for all commands
  work help dirs  # shows help for 'dirs' command
  work help trees # shows help for 'trees' command`,
	Call: func(x *bonzai.Cmd, args ...string) error {
		cmds := []*bonzai.Cmd{dirsCmd, treeCmd}
		if len(args) == 0 {
			x.Println("{{.Name}} - {{.Short}}")
			x.Println("")
			x.Println("USAGE:")
			x.Println("  {{.Usage}}")
			x.Println("")
			x.Println("COMMANDS:")
			for _, cmd := range cmds {
				fmt.Printf("  %-10s - %s\n", cmd.Name, cmd.Short)
			}
			fmt.Printf(
				"\nUse 'help [command]' for more information about a command.\n",
			)
		} else if len(args) == 1 {
			var matched bool
			for _, cmd := range cmds {
				if cmd.Name == args[0] {
					matched = true
					cmd.Println("{{.Name}} - {{.Short}}")
					cmd.Println("")
					cmd.Println("USAGE:")
					cmd.Println("  {{.Usage}}")
					cmd.Println("")
					cmd.Println("DESCRIPTION:")
					cmd.Println("  {{.Long}}")
				}
			}
			if !matched {
				return fmt.Errorf("unknown command %q", args[0])
			}
		} else {
			return fmt.Errorf("too many arguments")
		}
		return nil
	},
}

// dirsCmd lists work directories with optional compact output.
var dirsCmd = &bonzai.Cmd{
	Name:  `dirs`,
	Alias: `d`,
	Usage: `dirs`,
	Short: `list work directories`,
	Long: `
Lists local work directories, including Git worktree repositories.
Set SHORT=1 for a compact output format.

ENVIRONMENT VARIABLES
  SHORT|W_SHORT     Set SHORT=1 to display compact output
`,
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

// treeCmd lists Git worktrees with optional compact output.
var treeCmd = &bonzai.Cmd{
	Name:  `trees`,
	Alias: `t`,
	Usage: `work trees`,
	Short: `list Git worktrees`,
	Long: `
Lists Git worktrees within the specified directories.
Use SHORT=1 for a compact output format.

ENVIRONMENT VARIABLES
  SHORT     Set SHORT=1 to display compact output
  W_SHORT   Alternate way to set SHORT=1 for compact output
`,
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
