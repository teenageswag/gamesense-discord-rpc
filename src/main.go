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
	"1457019638331347008",
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

	if !process.DiscordRunning() {
		log.Info("Процесс discord не найден. Пробуем запустить вручную")
		if !process.LaunchDiscord() {
			return
		}
		// log.Info("Discord запущен, ожидаем полной инициализации...")
		// time.Sleep(8 * time.Second)

		log.Info("Ожидание запуска Discord IPC.")
		if !waitForDiscordPipe(60 * time.Second) {
			log.Error("Не удалось подключиться к Discord IPC (таймаут)")
			log.Warning("Попробуйте перезапустить программу через несколько секунд")
			return
		}
		log.Info("Ожидание готовности Discord RPC...")
		time.Sleep(3 * time.Second)

	} else {
		log.Success("Discord найден. Подключаемся к RPC.")
	}

	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err = client.Login(gs.AppId)
		if err == nil {
			break
		}

		if strings.Contains(err.Error(), "pipe is being closed") {
			log.Warning("Pipe ещё не готов, повторная попытка через 2 секунды.")
			time.Sleep(2 * time.Second)
			continue
		}

		log.Error("Не удалось запустить RPC: " + err.Error())
		return
	}

	if err != nil {
		log.Error("Не удалось подключиться после " + string(rune(maxRetries)) + " попыток")
		return
	}

	defer client.Logout()

	startTime := time.Now()
	err = rpc.UpdatePresence(gs, startTime)
	if err != nil {
		log.Error("Не удалось установить активность")
		return
	}

	log.Success("RPC успешно запущен!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rpc.UpdatePresence(gs, startTime)
		case <-sigChan:
			log.Info("Завершение работы RPC")
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
			log.Success("Discord IPC pipe обнаружен")
			return true
		}
		if attemptCount%5 == 0 {
			log.Info("Ожидание IPC pipe...")
		}
		time.Sleep(1 * time.Second)
	}
	return false
}
