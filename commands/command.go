package commands

import "log"

type CommandInterface interface {
	Execute([]string, *log.Logger)
}

type Command struct {
	Command     string
	Description string
	Interface   CommandInterface
}

type CommandManager struct {
	Commands map[string]*Command
}

func (c *CommandManager) AddCommand(command *Command) {
	c.Commands[command.Command] = command
}

func (c *CommandManager) ExecuteCommand(name string, parameter []string, l *log.Logger) {
	if command, ok := c.Commands[name]; ok {
		command.Interface.Execute(parameter, l)
	}
}

func NewCommand(command string, description string, execte CommandInterface) *Command {
	return &Command{
		Command:     command,
		Description: description,
		Interface:   execte,
	}
}

func (c *CommandManager) CheckCommand(name string) bool {
	if _, ok := c.Commands[name]; ok {
		return true
	}
	return false
}
