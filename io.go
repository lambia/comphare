package main

import (
	"io/ioutil"
	"encoding/hex"
)


/*
x := []byte("test123")
err := save("file.ext",x)
if err != nil {
    log.Fatal(err)
}

y, err := loadBytes("file.ext")
if err != nil {
    log.Fatal(err)
} else {
    fmt.Printf(string(y[:]))
}
*/

func save(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0600)
}

func load(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	fileContent := hex.EncodeToString(content)
	return fileContent, nil
}

func loadBytes(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}