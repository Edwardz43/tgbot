package main

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/crawl/ptt"
	"Edwardz43/tgbot/err"
	"Edwardz43/tgbot/log"
	"Edwardz43/tgbot/log/zaplogger"
	"Edwardz43/tgbot/message/from"
	"Edwardz43/tgbot/worker"
	"Edwardz43/tgbot/worker/rabbitmqworker"
	"regexp"
	"strings"
)

var logger log.Logger
var jobWorker worker.Worker
var failOnError = err.FailOnError

func main() {
	logger = zaplogger.GetInstance()
	jobWorker = rabbitmqworker.GetInstance(logger)
	go jobWorker.Do(CrawlPTT)
	select {}
	//serve()
}

// CrawlPTT crawls the target board from PTT
func CrawlPTT(arg ...interface{}) error {
	result := arg[0].(*from.Result)
	msg := strings.Split(result.Message.Text, "@")
	cmd := msg[0]
	target := msg[1]
	if target != config.GetBotID() {
		return nil
	}
	isCommand, err := regexp.MatchString(`^/C[a-z]+$`, cmd)
	failOnError(err, "error when regex tgbot message")

	if !isCommand {
		return nil
	}

	m := ptt.BoardMap

	var board string

	if value, ok := m[cmd]; ok {
		board = value
	} else {
		return nil
	}

	crawler := ptt.GetInstance(board)
	s := crawler.Get()

	c := &Command{
		ChatID:    result.Message.Chat.ID,
		Text:      s,
		ParseMode: "HTML",
	}

	return send(&c)
}
