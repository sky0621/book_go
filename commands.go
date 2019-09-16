package main

import (
	"fmt"
	"os"
)

func NewCommands(storeInfo StoreInfo) *Commands {
	_, err := os.Stat(storeInfo.GetName())
	if err != nil {
		NewClearCommand(storeInfo).Exec(nil)
	}
	return &Commands{
		commands: map[string]Command{
			"end":    NewEndCommand(),
			"help":   NewHelpCommand(),
			"clear":  NewClearCommand(storeInfo),
			"save":   NewSaveCommand(storeInfo),
			"get":    NewGetCommand(storeInfo),
			"remove": NewRemoveCommand(storeInfo),
			"list":   NewListCommand(storeInfo),
		},
	}
}

type Commands struct {
	commands map[string]Command
}

func (c *Commands) Exec(cmds []string) {
	if len(cmds) < 1 {
		fmt.Println("no target")
		return
	}

	cmd := c.commands[cmds[0]]
	if cmd == nil {
		fmt.Println("no target")
		return
	}

	if len(cmds) == 1 {
		cmd.Exec(nil)
		return
	}
	cmd.Exec(cmds[1:])
}
