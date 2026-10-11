package mail

import (
	"fmt"
	"jfcarr/unread-mail-count/models"
	"jfcarr/unread-mail-count/util"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// Iterate through account list, retrieving unread mail count for each.  Uses a goroutine for each check.
func CheckAccounts(accounts []models.Account, settings models.Settings) {
	results := make(chan models.Result, len(accounts))
	var wg sync.WaitGroup

	for _, account := range accounts {
		wg.Add(1)

		go func(a models.Account) {
			defer wg.Done()
			results <- getUnreadCount(a)
		}(account)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var summaryString string = ""

	for r := range results {
		if r.Err != nil {
			log.Printf("%s: check failed: %v", r.Account, r.Err)
			continue
		}

		foundUnread := false

		if r.Unread > 0 {
			if foundUnread == true {
				summaryString = summaryString + " "
			}
			summaryString = summaryString + fmt.Sprintf("%s(%d)", r.Username, r.Unread)
			fmt.Printf("%d : %s\n", r.Unread, r.Username)

			foundUnread = true
		}

		if !foundUnread {
			summaryString = "No unread mail"
			if settings.DisplayZeroCounts {
				fmt.Printf("%s\n", summaryString)
			}
		}

		if settings.GenerateSummary {
			os.WriteFile(util.QualifiedFile("unread-mail-count-summary.txt"), []byte(summaryString), 0644)
		}
	}
}

// Retrieve unread mail count for given account.
func getUnreadCount(a models.Account) models.Result {
	password := util.GetPasswordFromEnvironment(a.PasswordKey)

	dialer := &net.Dialer{Timeout: 5 * time.Second}
	c, err := imapclient.DialTLS(a.Server, &imapclient.Options{
		Dialer: dialer,
	})
	if err != nil {
		return models.Result{Account: a.Name, Username: a.Username, Err: err}
	}

	defer c.Close()

	if err := c.Login(a.Username, password).Wait(); err != nil {
		return models.Result{Account: a.Name, Username: a.Username, Err: err}
	}

	status, err := c.Status(a.Mailbox, &imap.StatusOptions{
		NumUnseen: true,
	}).Wait()
	if err != nil {
		return models.Result{Account: a.Name, Username: a.Username, Err: err}
	}

	if status.NumUnseen == nil {
		return models.Result{Account: a.Name, Username: a.Username, Err: fmt.Errorf("server did not return unseen count")}
	}

	return models.Result{Account: a.Name, Username: a.Username, Unread: *status.NumUnseen}
}
