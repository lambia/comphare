package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    
    //"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
    /*
    vars := mux.Vars(r)
    fileName := vars["fileName"]
	fmt.Fprintln(w, "File name:", fileName)
    */
    
    /*
	files := DumpFiles{
        DumpFile{Filename: "test.mfd", Content: []byte{'t', 'e', 's', 't'} },
        DumpFile{Filename: "file.ext", Content: []byte{'f', 'i', 'l', 'e'} },
    }
    */

    fileContent, err := load("files/milo-test.mfd")
    if err != nil {
        panic(err)
    }
    file := DumpFile{Filename: "Milo-Test", Content: fileContent}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(file); err != nil {
        panic(err)
    }
}