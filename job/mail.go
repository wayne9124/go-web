package job

import (
	"encoding/json"
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"github.com/pkg/errors"
	"net/smtp"
)

type MailStruct struct {
	To      []string `json:"to"`
	Message string   `json:"message"`
}

// Send an email
// The payload must be a JSON object compliant with the MailStruct structure.
func (Job) Mail(payload string) (bool, error) {
	var data MailStruct
	conf, _ := gwf.Configuration()

	auth := smtp.PlainAuth("", conf.Mail.Host, conf.Mail.Password, conf.Mail.Host)
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return false, errors.New("Cannot unmarshal payload")
	}

	server := fmt.Sprintf("%s:%d", conf.Mail.Host, conf.Mail.Port)
	if err := smtp.SendMail(server, auth, conf.Mail.From, data.To, []byte(data.Message)); err != nil {
		gwf.ProcessError(err)
	}

	return true, nil
}
