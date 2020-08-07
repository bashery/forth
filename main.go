package main

import (
	//"log"
	//"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	url := launcher.New().Headless(false).Launch()
	page := rod.New().ControlURL(url).Connect().Trace(true).Timeout(15 * time.Second).Page("https://google.com")

	page.Element(`input[name=q]`).WaitVisible().Input("chromedp").Press(input.Enter)
	// #rso > div:nth-child(1) > div > div.r > a > h3
	//page.Element(`#rso > div:nth-child(1) > div > div > div.r > a > h3`).Click()
	page.Element(`#rso > div:nth-child(1) > div > div > div.r > a > h3`).Click()
	time.Sleep(time.Second * 5)

}
