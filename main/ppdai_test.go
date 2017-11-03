package main

import (
	_ "cloan/config"
	"cloan/platforms"
	"cloan/structs"
	"encoding/json"
	"testing"
)

func TestPPDai_Crawl(t *testing.T) {
	var crawler platforms.Crawler
	crawler = new(PPDai)

	bills, err := crawler.Crawl(structs.CrawlParams{Phone: "13903415520", Password: "y2767018"})
	if err != nil {
		t.Error("", err)
	} else {
		jsonStr, _ := json.Marshal(bills)
		log.Notice("Crawler output: %s", string(jsonStr))
	}
}
