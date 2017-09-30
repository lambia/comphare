/*
 * Comphare - A fare-compare tool
 * by Luca Lambia - 2017
 *
 * ToDo:
 *	- check if net/http could be good for THIS case
 *	- fix run-all*
 *  - add dumpfile ID based on md5
 *  - remove panics and fatals
 *  - this is not a real webserver
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