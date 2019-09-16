package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewRemoveCommand(storeInfo StoreInfo) Command {
	return &removeCommand{storeInfo: storeInfo}
}

type removeCommand struct {
	storeInfo StoreInfo
}

func (c *removeCommand) Exec(args []string) {
	if len(args) != 1 {
		fmt.Printf("not valid: %#v\n", args)
		return
	}

	b, err := ioutil.ReadFile(c.storeInfo.GetName())
	if err != nil {
		fmt.Printf("error1: %s\n", err.Error())
		return
	}

	var d map[string]string
	err = json.Unmarshal(b, &d)
	if err != nil {
		fmt.Printf("error2: %s\n", err.Error())
		return
	}
	delete(d, args[0])

	res, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("error3: %s\n", err.Error())
		return
	}

	err = ioutil.WriteFile(c.storeInfo.GetName(), res, os.ModePerm)
	if err != nil {
		fmt.Printf("error4: %s\n", err.Error())
		return
	}
}
