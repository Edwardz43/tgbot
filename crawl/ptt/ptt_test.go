package ptt_test

import (
	"Edwardz43/tgbot/crawl/ptt"
	"log"
	"testing"
)

func TestGetPTTBeauty(t *testing.T) {
	crawler := ptt.Crawler{}
	s := crawler.Get()
	l := len(s)
	log.Println(l)
}
