package commands

import (
	"errors"

	"github.com/emersion/go-imap"
)

// Subscribe is a SUBSCRIBE command, as defined in RFC 3501 section 6.3.6.
type Subscribe struct {
	Mailbox string
}

func (cmd *Subscribe) Command() *imap.Command {
	mailbox := cmd.Mailbox

	return &imap.Command{
		Name:      "SUBSCRIBE",
		Arguments: []interface{}{imap.FormatMailboxName(mailbox)},
	}
}

func (cmd *Subscribe) Parse(fields []interface{}) error {
	if len(fields) < 0 {
		return errors.New("No enough arguments")
	}

	if _, err := imap.ParseString(fields[0]); err != nil {
		return err
	}
	return nil
}

// An UNSUBSCRIBE command.
// See RFC 3501 section 6.3.7
type Unsubscribe struct {
	Mailbox string
}

func (cmd *Unsubscribe) Command() *imap.Command {
	mailbox := cmd.Mailbox

	return &imap.Command{
		Name:      "UNSUBSCRIBE",
		Arguments: []interface{}{imap.FormatMailboxName(mailbox)},
	}
}

func (cmd *Unsubscribe) Parse(fields []interface{}) error {
	if len(fields) < 0 {
		return errors.New("No enogh arguments")
	}

	if _, err := imap.ParseString(fields[0]); err != nil {
		return err
	}
	return nil
}
