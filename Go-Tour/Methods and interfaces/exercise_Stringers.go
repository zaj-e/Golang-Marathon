package main

import "fmt"

type IPAddr [4]byte

func (self IPAddr) String() string {
	var ret string = ""
	for i := 0; i < len(self); i++ {
		if i != len(self)-1 {
			ret += fmt.Sprint(int(self[i])) +"."
		} else {
			ret += fmt.Sprint(int(self[i]))
		}
	}
	return ret
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
