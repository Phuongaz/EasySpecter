package commands

import "log"

type Help struct{}

func (Help) Execute(parameters []string, l *log.Logger) {
	for name, command := range GetCommandManager().Commands {
		l.Println(name + " - " + command.Description)
	}
}
