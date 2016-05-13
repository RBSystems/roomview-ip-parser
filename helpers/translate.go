package helpers

import (
	"errors"
	"net"
)

func TranslateToTP(h []Host) []Host { // Add 2 to the last value in the ipaddress
	toReturn := []Host{}

	for i := range h {
		ip := net.ParseIP(h[i].IPaddress)
		ip = ip.To4()
		if ip == nil {
			panic(errors.New("Bad IP"))
		}

		ip[3]++
		ip[3]++

		toReturn = append(toReturn, Host{Hostname: h[i].Hostname, IPaddress: ip.String()})
	}

	return toReturn
}
