package apollostats

import (
	"fmt"
	"sort"
)

type Command struct {
	Name   string
	Desc   string
	DoFunc func()
}

// The list of available commands.
var commands []*Command

// Implements the sort.interface for commands, based on the name.
type sortbyname []*Command

func (s sortbyname) Len() int {
	return len(s)
}

func (s sortbyname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortbyname) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Add a new command.
func AddCommand(name, desc string, dofunc func()) {
	c := &Command{name, desc, dofunc}
	commands = append(commands, c)
	sort.Sort(sortbyname(commands))
}

// Parse input and run matching command.
func RunCommand(input string) {
	for _, c := range commands {
		if input == c.Name {
			c.DoFunc()
			return
		}
	}
}

// Get a string of available commands.
func GetHelp() string {
	s := "Available commands:\n"
	for _, c := range commands {
		s += fmt.Sprintf("\t%s\t%s\n", c.Name, c.Desc)
	}
	return s
}
