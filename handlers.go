package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
	"encoding/json"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fileName := vars["fileName"] //is a map[string]string, returns "{value}"
    fileName = "files/"+fileName[1:len(fileName)-1]

	//ToDo: load() array of files
	fileContent, err := load(fileName)
	if err != nil {
		panic(err)
	}
	file := DumpFile{Filename: fileName, Content: fileContent}
	//ToDo: files:=DumpFiles{Dumpfile{...},...}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(file); err != nil {
		panic(err)
	}
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    /*
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
         */

}