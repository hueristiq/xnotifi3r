package slack

import (
	"net/url"

	"github.com/containrrr/shoutrrr"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
)

type Platform struct {
	*configuration.Slack
}

func New(conf *configuration.Slack) (platform *Platform, err error) {
	platform = &Platform{conf}

	return
}

func (platform *Platform) Send(message string) (err error) {
	url := &url.URL{
		Scheme: "slack",
		Path:   platform.Token + "@" + platform.ChannelID,
	}

	q := url.Query()

	if platform.Botname != "" {
		q.Set("botname", platform.Botname)
	}

	url.RawQuery = q.Encode()

	if err = shoutrrr.Send(url.String(), message); err != nil {
		return
	}

	return
}
