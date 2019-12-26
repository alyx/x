package mail

import "gopkg.in/gomail.v2"

// Mailer ...
type Mailer struct {
	Host string
	Port int
	User string
	Pass string
	From string
}

// Dialer builds a gomail.Dialer instance from
// a Mailer structure
func (m *Mailer) Dialer() *gomail.Dialer {
	d := gomail.NewDialer(m.Host, m.Port,
		m.User, m.Pass)

	return d
}

// SendText sends a text/plain-style email
func (m *Mailer) SendText(to string, subject string, content string) error {
	e := gomail.NewMessage()
	e.SetHeader("From", m.From)
	e.SetHeader("To", to)
	e.SetHeader("Subject", subject)
	e.SetBody("text/plain", content)

	d := m.Dialer()

	if err := d.DialAndSend(e); err != nil {
		return err
	}

	return nil
}

// Generate ...
func Generate(host string, port int, user string,
	pass string, from string) *Mailer {
	s := Mailer{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		From: from,
	}

	return &s
}
