package main

import (
	"encoding/json"
	"fmt"
)

func worldStatus() (err error) {

	const gateURL = "https://frontier.ffxiv.com/worldStatus/gate_status.json"
	const statusURL = "https://frontier.ffxiv.com/worldStatus/current_status.json"

	// Gate
	fmt.Print("Checking gate status ..")
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
		fmt.Println("Gate: Down")
	} else {
		fmt.Println(" OK")
	}

	// World status
	fmt.Print("Checking world status ..")
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
	for key, value := range statusMap {
		if value == 0 {
			fmt.Println(key, ": Down")
			maint = true
		}
	}
	if !maint {
		fmt.Println(" OK")
	}

	fmt.Println("")

	return
}
