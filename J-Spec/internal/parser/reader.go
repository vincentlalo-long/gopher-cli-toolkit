package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadJson(source io.Reader) (any, error) {
	var data any
	decoder := json.NewDecoder(source)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error , cannot decode")
	}
	return data, err
}

func OpenFile(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
