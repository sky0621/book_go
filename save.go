package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewSaveCommand(storeInfo StoreInfo) Command {
	return &saveCommand{storeInfo: storeInfo}
}

type saveCommand struct {
	storeInfo StoreInfo
}

func (c *saveCommand) Exec(args []string) {
	if len(args) != 2 {
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
	d[args[0]] = args[1]

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
