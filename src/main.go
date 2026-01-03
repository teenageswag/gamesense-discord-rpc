package main

import (
	"rpc/log"
	"rpc/process"
)

/*
	1. Проверяем запущен ли дс -> Готово
	2. Если не запущен, то запускаем автоматически
	3. Если запущен, то запускаем rpc
	4. Закрываем консоль и продолжаем работу в фоне
*/

func main() {
	// 1. Проверка запуска дс
	log.Info("Проверяем запущен ли дискорд...")
	if !process.DiscordRunning() {
		log.Error("Процесс не найден. Пробуем запустить вручную.")
		if !process.LaunchDiscord() {
			return
		}
	} else {
		log.Info("Discord уже запущен")
	}
}
