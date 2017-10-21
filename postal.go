//Package postal is a mock SMTP mailer
package postal

import (
	"net/smtp"
)

// MailRecord represents a single record of the SMTP transaction sent
type MailRecord struct {
	Host string
	Auth smtp.Auth
	From string
	To   []string
	Body []byte
}

// Postal represents a mocked SMTP mailer
type Postal struct {
	mailed      int
	mailRecords []MailRecord
	err         error
}

// New initializes and returns a new Postal object
func New() *Postal {
	return &Postal{}
}

// Mailed returns the number of messages mailed
func (p *Postal) Mailed() int {
	return p.mailed
}

// Mailer returns a mocked smtp mailer method capture and force actions for assertions
func (p *Postal) Mailer() func(string, smtp.Auth, string, []string, []byte) error {
	return func(host string, auth smtp.Auth, from string, to []string, body []byte) error {
		if p.err != nil {
			return p.err
		}

		record := MailRecord{
			Host: host,
			Auth: auth,
			From: from,
			To:   to,
			Body: body,
		}
		p.mailRecords = append(p.mailRecords, record)
		p.mailed++

		return nil
	}
}

// MailRecords returns the collection of mail records Postal has captured
func (p *Postal) MailRecords() []MailRecord {
	return p.mailRecords
}

// SetError sets the error the mailer is to return when called
func (p *Postal) SetError(err error) {
	p.err = err
}
