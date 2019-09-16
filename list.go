package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func NewListCommand(storeInfo StoreInfo) Command {
	return &listCommand{storeInfo: storeInfo}
}

type listCommand struct {
	storeInfo StoreInfo
}

func (c *listCommand) Exec(args []string) {
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

	println(`"key","value"`)
	for k, v := range d {
		println("\"" + k + "\",\"" + v + "\"")
	}

}
