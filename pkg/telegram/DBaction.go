package telegram

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)
type Article struct{
	Id uint16
	Task string
}
var posts = []Article{}
func (b* Bot) handleAddToDB(message *tgbotapi.Message) error {
	message.Text = strings.SplitAfterN(message.Text, "/add", 2)[1]
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if message.Text == ""{
		return err
	}
	insert, err := db.Query(fmt.Sprintf("INSERT INTO `tasks` (`task`, `priority`) VALUES('%s', '%d')", message.Text, 3))
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	return err
}
func (b* Bot) handleShowToDB(message *tgbotapi.Message) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res,err := db.Query("SELECT `id`, `task` FROM `tasks`")
	if err != nil {
		panic(err)
	}
	posts = []Article{}
	for res.Next(){
		var post Article
		err = res.Scan(&post.Id, &post.Task)
		if err != nil {
			panic(err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, post.Task)
		_, err = b.bot.Send(msg)
		if err !=nil{
			return err
		}
		posts = append(posts, post)
	}
	return err
}
func (b* Bot) handleRemoveToDB(message *tgbotapi.Message) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(message.Text)
	message.Text = strings.Split(message.Text, "remove ")[1]
	fmt.Println(message.Text)
	insert, err := db.Query(fmt.Sprintf("DELETE FROM `tasks` WHERE `task` = %s", message.Text))
	if err != nil {
		return err
	}
	defer insert.Close()
	return err
}