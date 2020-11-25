package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseJsonFile(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(v)
	if err != nil {
		return err
	}
	fmt.Println(v)
	return nil
}
