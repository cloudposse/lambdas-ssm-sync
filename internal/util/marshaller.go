package util

import (
	"encoding/json"
	"os"
)

func Marshal(inputEvent interface{}) ([]byte, error) {
	outputStream, err := json.Marshal(inputEvent)
	if err != nil {
		return nil, err
	}

	return outputStream, nil
}

func Unmarshal(inputStream []byte) (map[string]interface{}, error) {
	var outputEvent map[string]interface{}
	err := json.Unmarshal(inputStream, &outputEvent)
	if err != nil {
		return nil, err
	}

	return outputEvent, nil
}

func UnmarshalEvent[T interface{}](inputStream []byte) (T, error) {
	var outputEvent T
	err := json.Unmarshal(inputStream, &outputEvent)
	if err != nil {
		return *new(T), err
	}

	return outputEvent, nil
}

func UnmarshalFile[T interface{}](path string) (T, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return *new(T), err
	}

	return UnmarshalEvent[T](contents)
}
