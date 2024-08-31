package servies

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONFile(filePath string) ([]Title, error) {
    var persons []Title
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(data, &persons); err != nil {
        return nil, err
    }

    return persons, nil
}
func WriteJSONFile(filePath string, persons []Title) error {
    data, err := json.MarshalIndent(persons, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filePath, data, 0644)
}
type Title struct{
	Data []Question `json:"data"`
}
type Question struct{
	Name string `json:"name"`
	Answer string `json:"answer"`
}