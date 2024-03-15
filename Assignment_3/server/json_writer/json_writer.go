package json_writer

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
    Water  int    `json:"water"`
    Wind   int    `json:"wind"`
    Status string `json:"status"`
}


func WriteResponseToJSON(response Response) error {
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data successfully written to the JSON file")
	return nil
}