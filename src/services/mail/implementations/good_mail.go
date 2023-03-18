package mail

import "fmt"

type GoodMail struct{}

func (g GoodMail) Send(from string, to string, body string) {
	fmt.Printf("Email from %s sent to %s: %s", from, to, body)
}
