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
		log.Error("Не удалось запустить Discord. Сделайте это вручную и попробуйте снова!")
		return false
	}
	log.Success("Discord запущен.")
	return true
}
