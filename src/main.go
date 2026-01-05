package main

import (
	"os"
	"os/signal"
	"rpc/log"
	"rpc/process"
	"rpc/rpc"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var gs *rpc.RPCInfo = rpc.Initialize(
	"YOUR_APP_ID",           // Application ID
	"YOUR_STATE",            // State
	"YOUR_DETAILS",          // Details
	"YOUR_LARGE_IMAGE_KEY",  // Large Image Key
	"YOUR_LARGE_IMAGE_TEXT", // Large Image Text
	"YOUR_SMALL_IMAGE_KEY",  // Small Image Key
	"YOUR_SMALL_IMAGE_TEXT", // Small Image Text
	"YOUR_BUTTON_1_LABEL",   // Button 1 Label
	"YOUR_BUTTON_1_URL",     // Button 1 URL
)

func main() {
	// Check if Discord is running
	if !process.DiscordRunning() {
		log.Info("Discord process not found. Attempting to launch manually...")
		if !process.LaunchDiscord() {
			return
		}
		// Wait IPC pipe
		log.Info("Waiting for Discord IPC...")
		if !waitForDiscordPipe(60 * time.Second) {
			log.Error("Failed to connect to Discord IPC (Timeout).")
			log.Warning("Please try restarting the application.")
			return
		}
		// useless btw
		// log.Info("Discord IPC found. Initializing RPC...")
		// time.Sleep(3 * time.Second)

	} else {
		log.Success("Discord process found. Connecting to RPC...")
	}

	// Attempt login Discord RPC
	var err error
	maxRetries := 5
	for range maxRetries {
		err = client.Login(gs.AppId)
		if err == nil {
			break
		}

		if strings.Contains(err.Error(), "pipe is being closed") {
			log.Warning("Pipe not ready, retrying in 2 seconds...")
			time.Sleep(2 * time.Second)
			continue
		}

		log.Error("Failed to start RPC: " + err.Error())
		return
	}

	if err != nil {
		log.Error("Failed to connect after " + string(rune(maxRetries)) + " attempts")
		return
	}

	defer client.Logout()

	startTime := time.Now()
	err = rpc.UpdatePresence(gs, startTime)
	if err != nil {
		log.Error("Failed to set activity")
		return
	}

	log.Success("RPC started successfully!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rpc.UpdatePresence(gs, startTime)
		case <-sigChan:
			log.Info("Stopping RPC...")
			return
		}
	}

}

func waitForDiscordPipe(timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	attemptCount := 0
	for time.Now().Before(deadline) {
		attemptCount++
		_, err := os.Stat(`\\.\pipe\discord-ipc-0`)
		if err == nil {
			return true
		}

		if attemptCount%5 == 0 { // every 5 sec
			log.Info("Waiting for IPC pipe...")
		}
		time.Sleep(1 * time.Second)
	}
	return false
}
