package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(msgType int, msg string) error {
	webhookUrl := "https://hooks.slack.com/services/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"
	if msgType == 2 {
		webhookUrl = "https://hooks.slack.com/services/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"
	}
	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
