package main

import (
	"fmt"
	"os"
)

func NewEndCommand() Command {
	return &endCommand{}
}

type endCommand struct {
}

func (c *endCommand) Exec(args []string) {
	fmt.Println("End!")
	os.Exit(-1)
}
