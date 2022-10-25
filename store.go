package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func initStore() *UserStore {
	var store UserStore
	f, err := os.Open(storePath)
	if err != nil {
		panic("error opening store")
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error closing file")
		}
	}(f)

	data, err := io.ReadAll(f)

	err = json.Unmarshal(data, &store)
	if err != nil {
		panic("Could not parse store")
		return nil
	}

	return &store
}
