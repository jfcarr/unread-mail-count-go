# unread-mail-count-go

Retrieves unread mail count from one or more IMAP server accounts.

## Configuration

For this example configuration, we'll assume that you're checking a single Gmail account, username "johndoe@gmail.com", with a password of "mypassword12345".

For each account that you plan to check, add an environment variable containing the password of that account.  For our example, you'd add something like this:

```bash
export GM_JDOE="mypassword12345"
```

Make a copy of the example `unread-mail-count.json` file and name it `unread-mail-count-live.json`

Update the live file with the details of your account(s):

```json
[
    {
        "Name": "Personal",
        "Server": "imap.gmail.com:993",
        "Username": "johndoe@gmail.com",
        "PasswordKey": "GM_JDOE",
        "Mailbox": "INBOX"
    }
]
```

Build the executable:

```bash
go build
```

Run it:

```bash
./unread-mail-count -display-unread -summary
```

Output:

```
No unread mail
```

If you want to deploy the mail checker to your personal bin directory, you can use the included Makefile to simplify the process:

```bash
make deploy
```
