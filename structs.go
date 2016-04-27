package main

type configuration struct {
	Roomviewfolder    string
	RoomviewFileNames []string
	RoomRegex         string
	OutputFile        string
}

type host struct {
	IPaddress string
	Hostname  string
}
