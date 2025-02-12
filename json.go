package main

import (
	"encoding/json"
	"io"
)

// Десериализация из io.ReadCloser
func (request *MainRequest) Decode(body io.ReadCloser, result any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

// Десериализация из json в result
func FormatJson(input []byte, result any) error {
	err := json.Unmarshal(input, &result)
	if err != nil {
		return err
	}
	return nil
}

// Сериализация в json
func ParseJson(input any) []byte {
	bytes, err := json.Marshal(&input)
	if err != nil {
		return []byte{}
	}

	return bytes
}
