package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func writeToJSON(data Data) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("data.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}