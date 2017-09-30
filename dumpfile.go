package main

type DumpFile struct {
	Filename	string	`json:"filename"`
    //Byte		[]byte	`json:"contentB"`
    Content		string	`json:"content"`
}
type DumpFiles []DumpFile