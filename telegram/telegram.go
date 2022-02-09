package telegram

import (
	"TaskManager/database"
	"TaskManager/repo"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConnectTG(key string) {
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	DB := make(chan database.Quer)
	go database.ConnectTg(DB)
	conDB := <-DB
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		text := update.Message.Command()
		textText := update.Message.Text

		//user := conDB.GetLoginTg(update.Message.Chat.ID)
		userId := update.Message.Chat.ID
		switch text {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "all":
			user := conDB.GetLoginTg(userId)
			if user.Email != "" {

				objs := *conDB.GetAllObj()
				var teststring []string

				for i, _ := range objs {
					enter := `
________________________________________________
`
					text := fmt.Sprintf("Объект: %s , Дата начала: %s , Стоимость: %d рублей. %s", objs[i].Name, objs[i].Start, objs[i].PriceSum, enter)
					//	addLine := "Объект" + objs[i].Name + "  дата начала:" + objs[i].Start + "   стоимость:" + objs[i].PriceSum + enter
					teststring = append(teststring, text)

				}

				str2 := strings.Join(teststring, "")
				msg.Text = str2
			} else {
				msg.Text = "Нет авторизации"
			}
		default:

			if strings.Contains(textText, "add") {

				var object = repo.Object{}
				obj := strings.Fields(textText)

				if len(obj) == 3 {

					object.Name = obj[1]
					object.PriceSum, _ = strconv.Atoi(obj[2])
					_ = conDB.CreateObj(&object)
					msg.Text = "Создано -> " + "Объект : " + obj[1] + "  стоимость " + obj[2] + " рублей"
				} else {
					msg.Text = "Ошибка"
				}
			}

			if strings.Contains(textText, "create") {

				var user = repo.User{}
				usr := strings.Fields(textText)

				if len(usr) == 2 {

					user.Email = usr[1]
					user.Telegram = userId
					_ = conDB.UpdateUserTg(&user)
					msg.Text = "Создано -> "
				} else {
					msg.Text = "Ошибка"
				}
			}

			if strings.Contains(textText, "del") {

				obj := strings.Fields(textText)
				if len(obj) == 2 {
					_ = conDB.DeleteObjTg(obj[1])
					msg.Text = "Удалено -> " + "Объект : " + obj[1]
				} else {
					msg.Text = "Ошибка"
				}
			}

		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
