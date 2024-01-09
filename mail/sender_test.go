package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valrichter/go-basic-bank/util"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Test Email"
	content := `
		<h1>HELLO WORD</h1>
		<p>This is a test email from <a href="https://github.com/valrichter/go-basic-bank">go-basic-bank</a> </p>
	`
	to := []string{"email@example.com"}
	attachFile := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFile)
	require.NoError(t, err)
}
