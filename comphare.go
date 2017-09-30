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
import "net/http"
import "github.com/gorilla/mux"
import "encoding/json"

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/file/{fileName}", FileHandler)
	fmt.Printf("Comphare serving on\r\nhttp://localhost:80\r\n")
	log.Fatal(http.ListenAndServe(":80", router))

	/*
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
	*/
}

type DumpFile struct {
	Filename      string    `json:"filename"`
    Content       []byte	`json:"content"`
}
type DumpFiles []DumpFile

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fileName := vars["fileName"]
	fmt.Fprintln(w, "File name:", fileName)
	
	todos := DumpFiles{
        DumpFile{Filename: "test.mfd", Content: []byte{'t', 'e', 's', 't'} },
        DumpFile{Filename: "file.ext", Content: []byte{'f', 'i', 'l', 'e'} },
    }

    json.NewEncoder(w).Encode(todos)
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