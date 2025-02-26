package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"boilerplate-api/lib/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"google.golang.org/api/gmail/v1"
)

type EmailParams struct {
	To              string
	From            string
	SenderEmail     string
	SubjectData     string
	SubjectTemplate string
	BodyData        interface{}
	BodyTemplate    string
	Lang            string
}

type gLogger interface {
	Fatal(args ...interface{})
}

type GmailConfig struct {
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	hostURL      string
	logger       gLogger
}

type GmailService struct {
	*gmail.Service
	logger gLogger
}

func NewGmailService(gmailConfig GmailConfig) GmailService {
	ctx := context.Background()

	oauthConfig := oauth2.Config{
		ClientID:     gmailConfig.clientID,
		ClientSecret: gmailConfig.clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  gmailConfig.hostURL, // e.g: "http://localhost" or deployed API url
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
	}
	token := oauth2.Token{
		AccessToken:  gmailConfig.accessToken,
		RefreshToken: gmailConfig.refreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}
	var tokenSource = oauthConfig.TokenSource(ctx, &token)
	_service, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		gmailConfig.logger.Fatal("failed to receive gmail client", err.Error())
	}

	return GmailService{
		Service: _service,
		logger:  gmailConfig.logger,
	}
}

func (g GmailService) SendEmail(params EmailParams) (bool, error) {
	to := params.To
	from := params.From
	sender := params.SenderEmail
	emailBody, err := utils.ParseTemplate(params.BodyTemplate, params.BodyData)
	if err != nil {
		return false, errors.New("unable to parse email body template")
	}
	var msgString string
	emailTo := "To: " + to + "\r\n"
	msgString = emailTo
	subject := "Subject: " + params.SubjectData + "\n"
	msgString = msgString + subject
	msgString = msgString + "\n" + emailBody
	var msg []byte

	var _from string

	if from != "" && sender != "" {
		// sender should be email from which mail is being sent
		if params.Lang != "en" {
			encodedName := base64.StdEncoding.EncodeToString([]byte(from))
			_from = fmt.Sprintf("From: =?UTF-8?B?%s?= <%s>\r\n", encodedName, sender)
		} else {
			_from = fmt.Sprintf("From: \"%s\" <%s>\r\n", from, sender)
		}
	}

	if _from != "" {
		msgString = _from + msgString
	}

	if params.Lang != "en" {
		msgStringJP, _ := utils.ToISO2022JP(msgString)
		msg = msgStringJP
	} else {
		msg = []byte(msgString)
	}
	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString(msg),
	}
	_, err = g.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}
