# Postal
[![Build Status](https://travis-ci.org/gomicro/postal.svg)](https://travis-ci.org/gomicro/postal)
[![Coverage](http://gocover.io/_badge/github.com/gomicro/postal)](http://gocover.io/github.com/gomicro/postal)
[![GoDoc](https://godoc.org/github.com/gomicro/postal?status.png)](https://godoc.org/github.com/gomicro/postal)

Postal is a mocked SMTP mailer.

# Usage
Compare a mail record

```
func testMailing(t *testing.T) {
	host := "mailhost.com:25"
	auth := smtp.PlainAuth("", "username", "password", "mailhost.com")
	from := "dev@gomicro.io"
	to := []string{"foo@bar.com"}
	body := []byte{"hello world"}

	p := postal.New()
	SendMail := p.Mailer()

	err := SendMail(host, auth, from, to, body)
	if err != nil {
		t.Error(err.Error())
	}

	if p.Mailed() != 1 {
		t.Errorf("expected 1 mailed, got %v mailed", p.Mailed())
	}

	records := p.MailRecords()
	if len(records) != 1 {
		t.Errorf("expected 1 record, got %v record", len(records))
	}

	r := records[0]
	if r.Host != host {
		t.Errorf("expected %v, got %v", host, r.Host)
	}

	if r.From != from {
		t.Errorf("expected %v, got %v", from, r.From)
	}

	if r.To != to {
		t.Errorf("expected %v, got %v", to, r.To)
	}

	if r.Body != body {
		t.Errorf("expected %v, got %v", string(body), string(r.Body))
	}
}
```

Get an error from sending an email

```
func testMailing(t *testing.T) {
	auth := smtp.PlainAuth("", "username", "password", "mailhost.com")

	p := postal.New()
	p.SetError(fmt.Errorf("something's not quite right here"))
	SendMail := p.Mailer()

	err := SendMail("mailhost.com:25", auth, "dev@gomicro.io", []string{"foo@bar.com"}, []byte("Hello world"))

	if err == nil {
		t.Errorf("Expected error, and got nil")
	}
}
```
