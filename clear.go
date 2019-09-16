package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewClearCommand(storeInfo StoreInfo) Command {
	return &clearCommand{storeInfo: storeInfo}
}

type clearCommand struct {
	storeInfo StoreInfo
}

func (c *clearCommand) Exec(args []string) {
	_, err := os.Create(c.storeInfo.GetName())
	if err != nil {
		fmt.Printf("error1: %s\n", err.Error())
		return
	}

	d := make(map[string]string, 1)
	b, err := json.Marshal(&d)
	if err != nil {
		fmt.Printf("error2: %s\n", err.Error())
		return
	}

	err = ioutil.WriteFile(c.storeInfo.GetName(), b, os.ModePerm)
	if err != nil {
		fmt.Printf("error3: %s\n", err.Error())
		return
	}
}
