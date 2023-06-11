package common

import (
	"context"
	"fmt"
	"go-web-boilerplate/shared/config"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"github.com/sirupsen/logrus"
)

type MailerRequest struct {
	Email string
	Name  string
}

func (m *MailerRequest) Mailer(env *config.EnvConfig, log *logrus.Logger, token string) {
	var (
		ctx context.Context
	)

	cfg := sendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", env.SendinBlueApiKey)

	body := sendinblue.SendSmtpEmail{
		Sender: &sendinblue.SendSmtpEmailSender{
			Name:  "YourAppName",
			Email: "no-reply@yourappname.com",
		},
		To: []sendinblue.SendSmtpEmailTo{
			{
				Email: m.Email,
				Name:  m.Name,
			},
		},
		TextContent: fmt.Sprintf("reset your password here %s", env.ClientPasswordResetUrl+token),
		Subject:     "YourAppName Password Reset",
	}

	log.Infof("sending email with context: %s", body)

	mailer := sendinblue.NewAPIClient(cfg)
	_, _, err := mailer.TransactionalEmailsApi.SendTransacEmail(ctx, body)

	if err != nil {
		log.Errorf("error sending email: %s", err.Error())
	}
}