package helpers

type Configuration struct {
	AddressBooksDirectory string
	RoomviewAddressBooks  []string
	RoomRegex             string
	OutputFile            string
}

type Host struct {
	IPaddress string
	Hostname  string
}
