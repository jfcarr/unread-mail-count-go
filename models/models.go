package models

// Application settings.
type Settings struct {
	DisplayZeroCounts bool
	GenerateSummary   bool
}

// Account to check.
type Account struct {
	Name        string
	Server      string // e.g. "imap.example.com:993"
	Username    string
	PasswordKey string
	Mailbox     string // e.g. "INBOX"
}

// Results of an account check.
type Result struct {
	Account  string
	Username string
	Unread   uint32
	Err      error
}
