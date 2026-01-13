package process

import (
	"os"
	"os/exec"
	"path/filepath"
	"rpc/log"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func DiscordRunning() bool {
	processes, _ := process.Processes()
	for _, p := range processes {
		n, _ := p.Name()
		if strings.Contains(strings.ToLower(n), strings.ToLower("discord")) {
			return true
		}
	}
	return false
}

func LaunchDiscord() bool {
	localAppData := os.Getenv("LOCALAPPDATA")
	discordPath := filepath.Join(localAppData, "Discord", "Update.exe")
	cmd := exec.Command(discordPath, "--processStart", "Discord.exe")

	err := cmd.Start()
	if err != nil {
		log.Error("Failed to launch Discord. Please launch it manually and try again!")
		return false
	}
	log.Success("Discord launched successfully.")
	return true
}
