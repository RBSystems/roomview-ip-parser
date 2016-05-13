package main

import (
	"fmt"

	"github.com/byuoitav/roomview-ip-parser/helpers"
)

func main() {
	config := helpers.ImportConfig("./config.json")
	helpers.Maps = make(map[string][]helpers.Host)

	addresses := []helpers.Host{}
	for file := range config.RoomviewAddressBooks {
		hosts, err := helpers.ParseFile(config.AddressBooksDirectory+"/"+config.RoomviewAddressBooks[file], config)
		if err != nil {
			fmt.Println("Could not read file: " + err.Error())
			return
		}

		addresses = append(addresses, hosts...)
	}

	addresses = []helpers.Host{}

	for _, v := range helpers.Maps {
		addresses = append(addresses, v[0])
	}

	ips := helpers.TranslateToTP(addresses)
	helpers.OutputToJSON(ips)
	helpers.OutputToTxt(ips)
}
