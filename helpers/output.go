package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func OutputToJSON(toWrite interface{}) {
	b, err := json.Marshal(toWrite)

	err = ioutil.WriteFile("output.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func OutputToTxt(toWrite []Host) {
	f, err := os.Create("output.txt")
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
