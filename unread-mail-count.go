package main

import (
	"flag"
	"jfcarr/unread-mail-count/mail"
	"jfcarr/unread-mail-count/models"
	"jfcarr/unread-mail-count/util"
)

func main() {
	displayZeroCounts := flag.Bool("display-unread", false, "Print a message to the console if there are no unread messages")
	generateSummary := flag.Bool("summary", false, "Generate a summary text file 'unread-mail-count-summary.txt' containing unread count")
	flag.Parse()

	var settings = models.Settings{DisplayZeroCounts: *displayZeroCounts, GenerateSummary: *generateSummary}

	var accounts []models.Account = util.LoadAccounts()
	mail.CheckAccounts(accounts, settings)
}
