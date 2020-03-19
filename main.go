package main

import (
	"Edwardz43/tgbot/crawl/ptt"
	"Edwardz43/tgbot/err"
	"Edwardz43/tgbot/log"
	"Edwardz43/tgbot/log/zaplogger"
	"Edwardz43/tgbot/message/from"
	"Edwardz43/tgbot/worker"
	"Edwardz43/tgbot/worker/rabbitmqworker"
	"regexp"
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
	cmd := result.Message.Text

	isCommand, err := regexp.MatchString(`^![a-z]+$`, cmd)
	failOnError(err, "error when regex tgbot message")

	if !isCommand {
		return nil
	}

	var board string

	if value, ok := ptt.BoardMap[cmd]; ok {
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
