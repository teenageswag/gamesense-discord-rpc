package rpc

import (
	"rpc/log"
	"time"

	"github.com/hugolgst/rich-go/client"
)

/*
	1. Структура RPC -> Готова
	2. Функция инициализации RPC -> Готова
	3. Функция обновления RPC
*/

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

func Initialize(id string, state string, details string, largeImage string, largeText string, smallImage string, smallText string, buttonName string, buttonUrl string) *RPCInfo {
	if id == "" || id == "YOUR_APP_ID" {
		log.Error("Неверный ID приложения")
		return nil
	}
	return &RPCInfo{
		AppId:      id,
		State:      state,
		Details:    details,
		LargeImage: largeImage,
		LargeText:  largeText,
		SmallImage: smallImage,
		SmallText:  smallText,
		ButtonName: buttonName,
		ButtonUrl:  buttonUrl,
	}
}

func UpdatePresence(rpcInfo *RPCInfo, startTime time.Time) {
	activity := client.Activity{
		State:      rpcInfo.State,
		Details:    rpcInfo.Details,
		LargeImage: rpcInfo.LargeImage,
		LargeText:  rpcInfo.LargeText,
		// SmallImage: "rogue",
		// SmallText:  "Rogue - Level 100",

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
		// log.Error("Ошибка обновления активности")
	} else {
		// log.Info("Активность обновлена")
	}
}
