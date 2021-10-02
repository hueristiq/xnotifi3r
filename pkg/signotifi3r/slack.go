package signotifi3r

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type requestBody struct {
	Channel string `json:"channel,omitempty"`
	Text    string `json:"text,omitempty"`
}

type Slack struct {
	client     *fasthttp.Client
	webHookURL string
}

func (slack *Slack) send(message string) error {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.SetRequestURI(slack.webHookURL)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(fasthttp.MethodPost)

	body, err := json.Marshal(requestBody{
		Text:    message,
		Channel: "notifications",
	})
	if err != nil {
		return err
	}

	req.SetBody(body)

	if err := slack.client.Do(req, res); err != nil {
		return err
	}

	if res.StatusCode() != fasthttp.StatusOK {
		return nil
	}

	return nil
}
