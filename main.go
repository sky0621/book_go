package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	commands := NewCommands(NewStoreInfo("store.json"))

	println("Start!")
	for {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			cmds := strings.Split(s.Text(), " ")
			commands.Exec(cmds)
		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	}
}
