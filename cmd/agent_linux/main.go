package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
)

var URL string

type ResultJSON struct {
	Hostname  string   `json:"hostname"`
	Username  string   `json:"username"`
	Addresses []string `json:"addresses"`
}

func getLocalAddresses() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var addresses []string = make([]string, 0)

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				if v.IP.To4() == nil {
					continue
				}
				addresses = append(addresses, fmt.Sprintf("%v : %s", i.Name, v))

			case *net.IPNet:
				if v.IP.To4() == nil {
					continue
				}
				addresses = append(addresses, fmt.Sprintf("%v : %s", i.Name, v))
			}
		}
	}

	return addresses
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	return hostname
}

func getUsername() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}

	return user.Username
}

func main() {
	// Get local addresses
	addrs := getLocalAddresses()

	// Get hostname
	hostname := getHostname()

	// Get username
	username := getUsername()

	// Build json for sending to server
	result := ResultJSON{
		Hostname:  hostname,
		Username:  username,
		Addresses: addrs,
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return
	}

	// Send json to server
	_, err = http.Post(URL, "application/json", bytes.NewBuffer(jsonResult))
	if err != nil {
		return
	}
}
