package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hex21h/lp_pp/internal/color_cli"
)

func ParsingGetPageCount() {
	res, err := http.Get("https://store77.net/telefony_apple/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	v := doc.Selection
	text := doc.Find("head").Text()
	fmt.Println(text, v)
}

func main() {

	fmt.Println("LP(learning project) price parsing start..............................[", color_cli.ColorGreen, "ok", color_cli.ResetColor, "]")

	ParsingGetPageCount()

}
