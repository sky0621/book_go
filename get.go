package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func NewGetCommand(storeInfo StoreInfo) Command {
	return &getCommand{storeInfo: storeInfo}
}

type getCommand struct {
	storeInfo StoreInfo
}

func (c *getCommand) Exec(args []string) {
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
	fmt.Println(d[args[0]])
}
