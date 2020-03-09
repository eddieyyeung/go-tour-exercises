package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr IPAddr) String() string {
	s := make([]string, len(ipaddr))
	for i, v := range ipaddr {
		s[i] = strconv.Itoa(int(v))
	}
	return strings.Join(s, ",")
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
