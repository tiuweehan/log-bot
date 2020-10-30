package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	botApiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	rawUserIds := strings.Split(os.Getenv("TELEGRAM_USER_ID"), " ")
	host := os.Getenv("LOG_SENDER_HOST")

	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(botApiToken)
	if err != nil {
		log.Panic(err)
	}

	rawMessageBody, err := ioutil.ReadAll(os.Stdin)

	rawMessageText := string(rawMessageBody)

	if rawMessageText == "" {
		return
	}

	messageText := fmt.Sprintf("[%v]\n%v", host, rawMessageText)

	for _, rawUserId := range rawUserIds {
		userId, err := strconv.ParseInt(rawUserId, 10,  64)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		message := tgbotapi.NewMessage(userId, messageText)
		bot.Send(message)
	}
}
