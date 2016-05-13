package helpers

import (
	"bufio"
	"os"
	"regexp"
)

var Maps map[string][]Host

func ParseFile(fileLocation string, config Configuration) ([]Host, error) {
	file, err := os.Open(fileLocation)
	toReturn := []Host{}

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
			toReturn = append(toReturn, Host{Hostname: matches[m][1], IPaddress: matches[m][2]})
			_, has := Maps[matches[m][2]]
			if has {
				Maps[matches[m][2]] = append(Maps[matches[m][2]], Host{Hostname: matches[m][1], IPaddress: matches[m][2]})
			} else {
				Maps[matches[m][2]] = []Host{Host{Hostname: matches[m][1], IPaddress: matches[m][2]}}
			}
		}
	}

	return toReturn, nil
}
