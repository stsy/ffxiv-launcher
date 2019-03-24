package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Check FFXIV server status
func worldStatus() (err error) {

	// Status URLs
	const (
		gateURL   = "https://frontier.ffxiv.com/worldStatus/gate_status.json"
		statusURL = "https://frontier.ffxiv.com/worldStatus/current_status.json"
	)

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

// Auto-Login
func autoLogin() (start bool) {
	user, err := user.Load(userPath)
	if err != nil {
		return
	}

	client, err := client.Load(clientPath)
	if err != nil {
		return
	}

	// Check if autologin is enabled
	if user.Auth.Session.AutoLogin {
		fmt.Println("Auto-login enabled")
		fmt.Println(" If launch fails:")
		fmt.Println("  Set auto_login to false or remove your session ID from config")
		// Check if session ID is stored
		if user.Auth.Session.ID != "" {
			t, err := time.Parse(time.RFC850, user.Auth.Session.Date)
			if err != nil {
				log.Panic(err)
			}
			duration := time.Now().Sub(t)
			exipire := 6 - duration.Hours()/24

			// Check if TestSID has expoired, 6 days?
			if exipire <= 0 {
				fmt.Println(" Your session ID has expired, login again to renew")
				return
			}

			// Start game
			fmt.Printf(" Your session ID will expire in %.0f days", exipire)

			var args = []string{
				"DEV.TestSID=" + key, user.Auth.Session.ID,
				"DEV.UseSqPack=1",
				"DEV.DataPathType=1",
				"DEV.MaxEntitledExpansionID=" + client.Game.Expansion,
				"SYS.Region=3",
				"language=1",
			}

			// Check if DX11 is enabled
			bin := "ffxiv.exe"
			if client.Game.Dx11 {
				bin = "ffxiv_dx11.exe"
			}

			// Start game
			Start(client.Game.Path.Game+bin, args)
			start = true
			return

		}
		fmt.Println(" Your session ID will be stored at login")
		return

	} else if user.Auth.Session.ID != "" || user.Auth.Session.Date != "" {
		// Remove session ID form config
		fmt.Println("Auto-login disabled")
		user.Auth.Session.Date = ""
		user.Auth.Session.ID = ""
		user.Save()
		fmt.Println(" Removed session ID / timestamp form config")
		return
	}
	return
}
