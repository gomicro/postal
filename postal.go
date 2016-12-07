package postal

import (
	"net/smtp"
)

type mailer func(string, smtp.Auth, string, []string, []byte) error

type MailRecord struct {
	Host string
	Auth smtp.Auth
	From string
	To   []string
	Body []byte
}

type Postal struct {
	mailed      int
	mailRecords []MailRecord
	err         error
}

func New() *Postal {
	return &Postal{}
}

func (p *Postal) Mailed() int {
	return p.mailed
}

func (p *Postal) Mailer() mailer {
	return func(host string, auth smtp.Auth, from string, to []string, body []byte) error {
		record := MailRecord{
			Host: host,
			Auth: auth,
			From: from,
			To:   to,
			Body: body,
		}
		p.mailRecords = append(p.mailRecords, record)
		p.mailed++

		return p.err
	}
}

func (p *Postal) MailRecords() []MailRecord {
	return p.mailRecords
}

func (p *Postal) SetError(err error) {
	p.err = err
}
