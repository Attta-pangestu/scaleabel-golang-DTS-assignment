package json_writer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Response struct {
    Water  int    `json:"water"`
    Wind   int    `json:"wind"`
    Status string `json:"status"`
}


// WriteResponseToJSON is a function to write response data to a JSON file.
func WriteResponseToJSON(response Response) error {
	// Convert the response to JSON format
	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data successfully written to the JSON file")
	return nil
}