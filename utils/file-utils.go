package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}
func getSaveDirectory() string {
	return filepath.Join(getHomeDir(), ".local", "todo-cli")
}

func createSaveDir() string {
	appDir := getSaveDirectory()

	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		err := os.MkdirAll(appDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	return appDir
}
func GetSaveFilePath() string {
	return filepath.Join(createSaveDir(), "data.json")
}
func CreateDataFile() string {
	dataFile := GetSaveFilePath()
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		file, createErr := os.Create(dataFile)
		if createErr != nil {
			panic(createErr)
		}

		_, writeErr := file.WriteString("[]")
		if writeErr != nil {
			panic(writeErr)
		}

		closeErr := file.Close()
		if closeErr != nil {
			return ""
		}
	}
	return dataFile
}

func ReadFromDataFile[T any](dataFile string) (T, error) {
	file, readErr := os.ReadFile(dataFile)
	if readErr != nil {
		var zero T
		return zero, readErr
	}

	if len(file) == 0 {
		var zero T
		return zero, fmt.Errorf("empty JSON input")
	}

	var data T
	unmarshalErr := json.Unmarshal(file, &data)
	if unmarshalErr != nil {
		var zero T
		return zero, unmarshalErr
	}
	return data, nil
}

func AppendDataToFile[S []any, T any](dataFile string, newData T) (S, error) {
	data, readErr := ReadFromDataFile[S](dataFile)
	if readErr != nil {
		return data, readErr
	}
	data = append(data, newData)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	err = os.WriteFile(dataFile, dataBytes, 0777)
	if err != nil {
		fmt.Println(err)

		return data, err
	}
	return data, nil
}

func RemoveAllDataFromFile(dataFile string) {
	_, err := os.ReadFile(dataFile)
	if os.IsNotExist(err) {
		log.Fatal("no data exist")
	}
	removeErr := os.RemoveAll(dataFile)
	if removeErr != nil {
		return
	}
	CreateDataFile()
	println("all data cleared")
}
