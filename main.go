/*
 * Comphare - A fare-compare tool
 * by Luca Lambia - 2017
 *
 * ToDo:
 * - file struct returns content as array of chars
 * - new struct returns list of files in a directory [0, name, ext],[1, asd, exe],
 * - create bootstrap frontend with variables
 * - set global variables for config (version, dump folder/ext, port)
 * - check if net/http could be better of MUXing for THIS case
 * - remove panics and fatals
 * -- CLI version
 * - Add a db
 * -- Add tag array to DumpFile struct
 * -- Add dumpfile ID based on md5
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    router := NewRouter()
	fmt.Printf("Comphare serving on\r\nhttp://localhost:80\r\n")

	log.Fatal(http.ListenAndServe(":80", router))
}