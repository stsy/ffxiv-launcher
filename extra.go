package main

import (
	"encoding/json"
	"fmt"
)

// Check FFXIV server status
func worldStatus() (err error) {

	const gateURL = "https://frontier.ffxiv.com/worldStatus/gate_status.json"
	const statusURL = "https://frontier.ffxiv.com/worldStatus/current_status.json"

	fmt.Println("Checking server status:")
	// Gate status
	fmt.Print(" Gate: ")
	jsonGate, err := DownloadString(gateURL)
	if err != nil {
		return
	}

	gateMap := make(map[string]int)
	err = json.Unmarshal([]byte(jsonGate), &gateMap)
	if err != nil {
		return
	}

	if gateMap["status"] == 0 {
		fmt.Println("Maintenance")
	} else {
		fmt.Println("OK")
	}

	// Server status
	fmt.Print(" Servers: ")
	jsonStatus, err := DownloadString(statusURL)
	if err != nil {
		return
	}

	statusMap := make(map[string]int)
	err = json.Unmarshal([]byte(jsonStatus), &statusMap)
	if err != nil {
		return
	}

	var maint bool
	for server, status := range statusMap {
		if status == 0 {
			if !maint {
				fmt.Println("Maintenance")
			}
			// Print servers that are down
			fmt.Println(" -", server)
			maint = true
		}
	}
	if !maint {
		fmt.Println("OK")
	}

	fmt.Println("")
	return
}
