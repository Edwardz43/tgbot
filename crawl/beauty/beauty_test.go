package beauty_test

import (
	"Edwardz43/tgbot/crawl/beauty"
	"log"
	"testing"
)

func TestGetPTTBeauty(t *testing.T) {
	crawler := beauty.Crawler{}
	s := crawler.Get()
	l := len(s)
	log.Println(l)
}
