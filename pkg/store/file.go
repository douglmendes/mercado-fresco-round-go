package store

import (
	"encoding/json"
	"log"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

func New(fileName string) Store {
	return &FileStore{fileName}
}

type FileStore struct {
	FileName string
}

func (f *FileStore) Write(data interface{}) error {
	jsonContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("failed to parse memory data into JSON on store.Write")
		log.Println(err)
		return err
	}

	return os.WriteFile(f.FileName, jsonContent, 0644)
}

func (f *FileStore) Read(data interface{}) error {
	jsonContent, err := os.ReadFile(f.FileName)
	if err != nil {
		log.Println("failed to read file on store.Read")
		log.Println(err)
		return err
	}

	return json.Unmarshal(jsonContent, &data)
}
