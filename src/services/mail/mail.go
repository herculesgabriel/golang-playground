package mail

type Mail interface {
	Send(from string, to string, body string)
}
