package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"os"
	"regexp"
)

var maps map[string][]host

func parseFile(f string, config configuration) ([]host, error) {
	file, err := os.Open(f)
	toReturn := []host{}

	if err != nil {
		return toReturn, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(config.RoomRegex)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, 500)

		for m := range matches {
			toReturn = append(toReturn, host{Hostname: matches[m][1], IPaddress: matches[m][2]})
			_, has := maps[matches[m][2]]
			if has {
				maps[matches[m][2]] = append(maps[matches[m][2]], host{Hostname: matches[m][1], IPaddress: matches[m][2]})
			} else {
				maps[matches[m][2]] = []host{host{Hostname: matches[m][1], IPaddress: matches[m][2]}}
			}
		}
	}

	return toReturn, nil
}

func importConfig(file string) configuration {
	b, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var config configuration

	err = json.Unmarshal(b, &config)

	if err != nil {
		panic(err)
	}

	return config
}

func outputToJSONFile(toWrite interface{}, file string) {
	b, err := json.Marshal(toWrite)

	err = ioutil.WriteFile(file+".json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func outputToTxt(toWrite []host, file string) {
	f, err := os.Create(file + ".txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := range toWrite {
		str := ""
		str = str + toWrite[i].Hostname + "\t\t"
		str = str + toWrite[i].IPaddress + "\n"
		f.WriteString(str)
	}
}

//Add two to the last value in the ipaddress.
func translateToTP(h []host) []host {
	toReturn := []host{}

	for i := range h {
		ip := net.ParseIP(h[i].IPaddress)
		ip = ip.To4()
		if ip == nil {
			panic(errors.New("Bad IP"))
		}

		//ip = ip.Mask(ip.DefaultMask())
		ip[3]++
		ip[3]++

		toReturn = append(toReturn, host{Hostname: h[i].Hostname, IPaddress: ip.String()})
	}

	return toReturn
}

func main() {
	config := importConfig("./config.json")
	maps = make(map[string][]host)

	toReturn := []host{}
	for f := range config.RoomviewFileNames {
		hosts, err := parseFile(config.Roomviewfolder+"/"+config.RoomviewFileNames[f], config)
		if err != nil {
			panic(err)
		}

		toReturn = append(toReturn, hosts...)
	}

	toReturn = []host{}

	for _, v := range maps {
		toReturn = append(toReturn, v[0])
	}
	outputToJSONFile(toReturn, config.OutputFile)
	outputToTxt(toReturn, config.OutputFile)
	ips := translateToTP(toReturn)
	outputToTxt(ips, config.OutputFile+"IPs")
}
