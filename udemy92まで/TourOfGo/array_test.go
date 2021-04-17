package main

import "fmt"

type IPAddr [4]byte

type Array [2]byte

// TODO: Add a "String() string" method to IPAddr.
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func (arry Array) String() string {
	return fmt.Sprintf("%d += %d", arry[0], arry[1])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	arry := map[string]Array{
		"first": {4, 5},
	}
	fmt.Println(arry)
	fmt.Println(hosts, "\n---------------------------")
	// IPAddr = hosts["IPAddr"]
	// fmt.Println(IPAddr)
	// for ip := range hosts
}
