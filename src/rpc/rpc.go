package rpc

import (
	"os"
	"rpc/log"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type RPCInfo struct {
	AppId      string
	State      string
	Details    string
	LargeImage string
	LargeText  string
	SmallImage string
	SmallText  string
	ButtonName string
	ButtonUrl  string
}

func validate(val, placeholder string) string {
	if val == placeholder {
		return ""
	}
	return val
}

func Initialize(id string, state string, details string, largeImage string, largeText string, smallImage string, smallText string, buttonName string, buttonUrl string) *RPCInfo {
	if id == "" || id == "YOUR_APP_ID" {
		panic("Invalid App ID")
	}

	return &RPCInfo{
		AppId:      id,
		State:      validate(state, "YOUR_STATE"),
		Details:    validate(details, "YOUR_DETAILS"),
		LargeImage: validate(largeImage, "YOUR_LARGE_IMAGE_KEY"),
		LargeText:  validate(largeText, "YOUR_LARGE_IMAGE_TEXT"),
		SmallImage: validate(smallImage, "YOUR_SMALL_IMAGE_KEY"),
		SmallText:  validate(smallText, "YOUR_SMALL_IMAGE_TEXT"),
		ButtonName: validate(buttonName, "YOUR_BUTTON_1_LABEL"),
		ButtonUrl:  validate(buttonUrl, "YOUR_BUTTON_1_URL"),
	}
}

func UpdatePresence(rpcInfo *RPCInfo, startTime time.Time) error {
	activity := client.Activity{
		State:      rpcInfo.State,
		Details:    rpcInfo.Details,
		LargeImage: rpcInfo.LargeImage,
		LargeText:  rpcInfo.LargeText,

		Timestamps: &client.Timestamps{
			Start: &startTime,
		},

		Buttons: []*client.Button{
			{
				Label: rpcInfo.ButtonName,
				Url:   rpcInfo.ButtonUrl,
			},
		},
	}

	err := client.SetActivity(activity)
	if err != nil {
		log.Error("Error updating activity: " + err.Error())
		return err
	}
	return nil
}

func WaitDiscordPipe(timeout time.Duration) bool {
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
