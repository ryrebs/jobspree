package main

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func runOnce() {
	// Run to download playwright browsers.
	err := playwright.Install()
	log.Print((err))
}

type JobData struct {
	Link         string
	Title        string
	Requirements string
	Details      string
}

type JobApplyProcess interface {
	scrapeData(brw playwright.Browser) (interface{}, error)
	isTargetJob() bool
	createCover() string
	sendApplication()
}

func main() {
	pw, err := playwright.Run()
	options := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	}
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(options)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://mynimo.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

}
