package port

import (
	"net"
	"strconv"
	"time"
)

// ScanResult is
type ScanResult struct {
	Port    string
	State   string
	Service string
}

// ScanPort is
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

// InitialScan is
func InitialScan(hostname string) []ScanResult {

	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}

// WideScan is
func WideScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}
