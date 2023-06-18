package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var addr strings.Builder
	for i, v := range ip {
		addr.WriteString(strconv.Itoa(int(v)))
		if i != len(ip)-1 {	
			addr.WriteString(".")
		}
	}
	return addr.String()
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
