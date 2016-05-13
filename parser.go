package main

import (
	"fmt"

	"github.com/byuoitav/roomview-ip-parser/helpers"
)

func main() {
	config := helpers.ImportConfig("./config.json")
	helpers.Maps = make(map[string][]helpers.Host)

	toReturn := []helpers.Host{}
	for file := range config.RoomviewAddressBooks {
		hosts, err := helpers.ParseFile(config.AddressBooksDirectory+"/"+config.RoomviewAddressBooks[file], config)
		if err != nil {
			fmt.Println("Could not read file: " + err.Error())
			return
		}

		toReturn = append(toReturn, hosts...)
	}

	toReturn = []helpers.Host{}

	for _, v := range helpers.Maps {
		toReturn = append(toReturn, v[0])
	}

	// helpers.OutputToJSON(toReturn)
	// helpers.OutputToTxt(toReturn)

	ips := helpers.TranslateToTP(toReturn)
	helpers.OutputToJSON(ips)
	helpers.OutputToTxt(ips)
}
