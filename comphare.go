/*
 * Comphare - A fare-compare tool
 * by Luca Lambia - 2017
 *
 * ToDo:
 *
 */

package main

import "fmt"
import "io/ioutil"
import "log"

func main() {
	//fmt.Printf("hello, world\n")

	x := []byte("test123")
	err := save("file.ext",x)
	if err != nil {
		log.Fatal(err)
	}
	
	y, err := load("file.ext")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf(string(y[:]))		
	}
}

func save(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0600)
}

func load(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}