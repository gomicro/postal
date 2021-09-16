package postal

import (
	"fmt"
	"net"
	"net/smtp"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPostal(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Postal", func() {
		var postal *Postal

		g.BeforeEach(func() {
			postal = New()
		})

		g.It("should return the number of items mailed", func() {
			host := "localhost"
			port := "25"
			user := "malreynolds@serenity.com"
			pass := "patience"
			auth := smtp.PlainAuth("", user, pass, host)
			to := []string{"wash@sereity.com"}
			body := []byte("Take us out of the world, Wash. We got us some crime to be done.")

			m := postal.Mailer()
			err := m(net.JoinHostPort(host, port), auth, user, to, body)
			Expect(err).NotTo(HaveOccurred())
			Expect(postal.Mailed()).To(Equal(1))

			err = m(net.JoinHostPort(host, port), auth, user, to, body)
			Expect(err).NotTo(HaveOccurred())
			Expect(postal.Mailed()).To(Equal(2))
		})

		g.It("should record mail records", func() {
			host := "localhost"
			port := "25"
			user := "malreynolds@serenity.com"
			pass := "patience"
			auth := smtp.PlainAuth("", user, pass, host)
			to := []string{"wash@sereity.com"}
			body := []byte("Take us out of the world, Wash. We got us some crime to be done.")

			m := postal.Mailer()
			err := m(net.JoinHostPort(host, port), auth, user, to, body)
			Expect(err).To(BeNil())
			Expect(postal.Mailed()).To(Equal(1))

			err = m(net.JoinHostPort("gmail.com", port), auth, user, to, body)
			Expect(err).To(BeNil())
			Expect(postal.Mailed()).To(Equal(2))

			records := postal.MailRecords()
			Expect(len(records)).To(Equal(2))

			Expect(records[0].Host).To(Equal(net.JoinHostPort(host, port)))
			Expect(records[0].Auth).To(Equal(auth))
			Expect(records[0].From).To(Equal(user))
			Expect(records[0].To).To(Equal(to))
			Expect(records[0].Body).To(Equal(body))

			Expect(records[1].Host).To(Equal(net.JoinHostPort("gmail.com", port)))
			Expect(records[1].Auth).To(Equal(auth))
			Expect(records[1].From).To(Equal(user))
			Expect(records[1].To).To(Equal(to))
			Expect(records[1].Body).To(Equal(body))
		})

		g.It("should return an error when instructed to", func() {
			auth := smtp.PlainAuth("", "user", "pass", "gmail.com")
			m := postal.Mailer()

			err := m(net.JoinHostPort("gmail.com", "25"), auth, "Wash", []string{"Dino"}, []byte("This is a fertile land and we will thrive."))
			Expect(err).To(BeNil())

			expectedErr := fmt.Errorf("Curse your sudden but inevitable betrayal!")
			postal.SetError(expectedErr)
			err = m(net.JoinHostPort("gmail.com", "25"), auth, "Wash", []string{"Dino"}, []byte("This is a fertile land and we will thrive."))
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(expectedErr.Error()))
		})
	})
}
