package main

import (
	"fmt"

	"github.com/Ophois47/Go-Port-Scanner/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)

	//widescanresults := port.WideScan("localhost")
	//fmt.Println(widescanresults)
}
