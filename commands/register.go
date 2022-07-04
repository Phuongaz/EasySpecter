package commands

import (
	"log"
)

var CommandManagerInstance *CommandManager

func RegisterCommands(log *log.Logger) {
	CommandManagerInstance = &CommandManager{
		Commands: map[string]*Command{},
	}
	CommandManagerInstance.AddCommand(NewCommand("help", "Show help", Help{}))
	CommandManagerInstance.AddCommand(NewCommand("join", "Join a specter", Join{}))
	CommandManagerInstance.AddCommand(NewCommand("xbox", "Xbox Login", Xbox{}))
	CommandManagerInstance.AddCommand(NewCommand("specter", "Specter commands", Specter{}))
}

func GetCommandManager() *CommandManager {
	return CommandManagerInstance
}
