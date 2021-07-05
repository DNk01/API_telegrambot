package main

import (
	"github.com/DNk01/API_telegrambot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main(){

	bot, err := tgbotapi.NewBotAPI("/11:yourtoken")
	if err != nil{
		log.Fatal(err)
	}
	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil{
		log.Fatal()
	}
}
