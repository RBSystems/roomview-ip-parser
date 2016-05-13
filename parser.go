package main

import "github.com/byuoitav/roomview-ip-parser/helpers"

func main() {
	config := helpers.ImportConfig("./config.json")
	helpers.Maps = make(map[string][]helpers.Host)

	toReturn := []helpers.Host{}
	for f := range config.RoomviewFileNames {
		hosts, err := helpers.ParseFile(config.Roomviewfolder+"/"+config.RoomviewFileNames[f], config)
		if err != nil {
			panic(err)
		}

		toReturn = append(toReturn, hosts...)
	}

	toReturn = []helpers.Host{}

	for _, v := range helpers.Maps {
		toReturn = append(toReturn, v[0])
	}

	helpers.OutputToJSON(toReturn, config.OutputFile)
	helpers.OutputToTxt(toReturn, config.OutputFile)

	ips := helpers.TranslateToTP(toReturn)
	helpers.OutputToTxt(ips, config.OutputFile+"IPs")
}
