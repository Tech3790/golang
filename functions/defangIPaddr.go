package main

import (
	"strings"
)

// DefangIPaddr is a function that returns a defanged form of a given ip address
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", len(address))
}
