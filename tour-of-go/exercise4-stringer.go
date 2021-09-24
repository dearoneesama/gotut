package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v * IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", v[0], v[1], v[2], v[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		newip := interface{}(ip).(IPAddr)
		fmt.Printf("%v: %v\n", name, newip)
	}
}
