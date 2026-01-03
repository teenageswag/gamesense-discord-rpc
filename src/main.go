package main

import (
	"os"
	"os/signal"
	"rpc/log"
	"rpc/process"
	"rpc/rpc"
	"time"

	"github.com/hugolgst/rich-go/client"
)

/*
	1. Проверяем запущен ли дс -> Готово
	2. Если не запущен, то запускаем автоматически
	3. Если запущен, то запускаем rpc
	4. Закрываем консоль и продолжаем работу в фоне
*/

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

	log.Info("ПРоверяем запущен ли discord")

	if !process.DiscordRunning() {
		log.Info("Процесс не найден. Пробуем запустить вручную")
		if !process.LaunchDiscord() {
			return
		}
		log.Success("Процесс найден.")
		log.Info("Запускаем RPC")
	}

	err := client.Login(gs.AppId)
	if err != nil {
		log.Error("Не удалось запустить RPC")
		return
	}
	defer client.Logout()

	startTime := time.Now()
	rpc.UpdatePresence(gs, startTime)

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
