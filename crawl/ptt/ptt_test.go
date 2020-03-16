package ptt_test

import (
	"Edwardz43/tgbot/crawl/ptt"
	"log"
	"testing"
)

func TestGetPTTCrawl(t *testing.T) {
	crawler := ptt.GetInstance("Lifeismoney")
	s := crawler.Get()
	l := len(s)
	log.Println(l)
}
