package main

import (
	"fmt"
	"gamesense-rpc/gamesense"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var gs = gamesense.NewGamesense(
	"YOUR_APP_ID",
	"",
	"Get Good - Get Gamesense",
	"gs_logo640",
	"gamesense.pub",
	"",
	"",
	"Get Gamesense",
	"https://gamesense.pub",
)

func main() {
	err := client.Login(gs.AppId)
	if err != nil {
		log.Fatal("[ERROR] Discord connection error:", err)
	}
	defer client.Logout()

	startTime := time.Now()
	UpdatePresence(startTime)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			UpdatePresence(startTime)

		case <-sigChan:
			fmt.Println("[INFO] Discord RPC disabled")
			return
		}
	}
}

func UpdatePresence(startTime time.Time) {
	activity := client.Activity{
		State:      gs.State,
		Details:    gs.Details,
		LargeImage: gs.LargeImage,
		LargeText:  gs.LargeText,
		// SmallImage: "rogue",
		// SmallText:  "Rogue - Level 100",

		Timestamps: &client.Timestamps{
			Start: &startTime,
		},

		Buttons: []*client.Button{
			{
				Label: gs.ButtonName,
				Url:   gs.ButtonUrl,
			},
		},
	}

	err := client.SetActivity(activity)
	if err != nil {
		log.Println("[ERROR] Error updating activity:", err)
	} else {
		fmt.Println("[INFO] Activity updated:", activity.State)
	}
}
