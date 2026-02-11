package main

import "fmt"

type LogW struct {
	IP        string
	HTTPCode  int
	Timestamp string
}

func main() {
	logs := []LogW{
		{"192.168.1.1", 200, "2023-01-01"},
		{"192.168.1.2", 404, "2023-02-01"},
		{"192.168.1.3", 500, "2023-03-01"},
		{"192.168.1.4", 302, "2023-04-01"},
		{"192.168.1.5", 403, "2023-05-01"},
	}

	errorLogs := filterErrorLogs(logs)

	for _, log := range errorLogs {
		fmt.Println("IP:", log.IP, "Code:", log.HTTPCode, "Time:", log.Timestamp)
	}
}

func filterErrorLogs(logs []LogW) []LogW {
	var result []LogW
	for _, log := range logs {
		if log.HTTPCode >= 400 && log.HTTPCode < 600 {
			result = append(result, log)
		}
	}
	return result
}
