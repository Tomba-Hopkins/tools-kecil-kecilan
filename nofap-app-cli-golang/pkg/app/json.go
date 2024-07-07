package app

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFileData(filepath string) (*NoFapData, error) {

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var daysData NoFapData

	err = json.Unmarshal(file, &daysData)
	if err != nil {
		return nil, err
	}

	return &daysData, nil

}


func WriteFileData(filepath string, data *NoFapData) error {
	file, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, file, 0644)
}