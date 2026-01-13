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
	"YOUR_APP_ID",           // Insert your Application ID
	"YOUR_STATE",            // Insert your State or leave empty
	"YOUR_DETAILS",          // Insert your Details or leave empty
	"YOUR_LARGE_IMAGE_KEY",  // Insert your Large Image Key or leave empty
	"YOUR_LARGE_IMAGE_TEXT", // Insert your Large Image Text or leave empty
	"YOUR_SMALL_IMAGE_KEY",  // Insert your Small Image Key or leave empty
	"YOUR_SMALL_IMAGE_TEXT", // Insert your Small Image Text or leave empty
	"YOUR_BUTTON_1_LABEL",   // Insert your Button 1 Label or leave empty
	"YOUR_BUTTON_1_URL",     // Insert your Button 1 URL or leave empty
)

func main() {
	if !process.DiscordRunning() {
		log.Info("Discord process not found. Attempting to launch manually...")
		if !process.LaunchDiscord() {
			return
		}
		log.Info("Waiting for Discord IPC...")
		if !rpc.WaitDiscordPipe(60 * time.Second) {
			log.Error("Failed to connect to Discord IPC (Timeout).")
			log.Warning("Please try restarting the application.")
			return
		}

	} else {
		log.Success("Discord process found. Connecting to RPC...")
	}

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
