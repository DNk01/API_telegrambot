package telegram

import (
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)
const(
	commandStart = "start"
	commandAdd = "add"
	commandShow = "show"
	commandRemove = "remove"
)

func (b* Bot) handleCommand(message *tgbotapi.Message) error{
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAdd:
		return b.handleAddCommand(message)
	case commandShow:
		return b.handleShowCommand(message)
	case commandRemove:
		return b.handleRemoveCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b* Bot) handleMessage(message *tgbotapi.Message){
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
func (b* Bot) handleUnknownCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Комманда не была распознана")
	_, err := b.bot.Send(msg)
	return err
}
func (b* Bot) handleStartCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Запускаем")
	_, err := b.bot.Send(msg)
	return err
}
func (b* Bot) handleAddCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Добавляем запись")
	_, err := b.bot.Send(msg)
	if err == nil {
		err = b.handleAddToDB(message)
	}
	return err
}
func (b* Bot) handleShowCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выводим текущие дела")
	_, err := b.bot.Send(msg)
	if err == nil{
		err = b.handleShowToDB(message)
	}
	return err
}
func (b* Bot) handleRemoveCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Удаление элемента с выбранным id")
	_, err := b.bot.Send(msg)
	if err == nil{
		err = b.handleRemoveToDB(message)
	}
	return err
}
