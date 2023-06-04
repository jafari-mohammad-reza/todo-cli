package utils

import (
	"encoding/json"
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

func CreateDataFile() string {
	dataFile := filepath.Join(createSaveDir(), "data.json")
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		file, createErr := os.Create(dataFile)
		if createErr != nil {
			panic(createErr)
		}
		closeErr := file.Close()
		if closeErr != nil {
			return ""
		}
	}
	return dataFile
}
func ReadFromDataFile[T []any](dataFile string) T {
	file, readErr := os.ReadFile(dataFile)
	if readErr != nil {
		panic(readErr)
	}
	var data T
	unmarshalErr := json.Unmarshal(file, &data)
	if unmarshalErr != nil {
		return nil
	}
	return data
}
func AppendDataToFile[T any](dataFile string, newData T) []any {
	data := ReadFromDataFile(dataFile)
	data = append(data, newData)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dataFile, dataBytes, 0644)
	if err != nil {
		panic(err)
	}
	return data
}
