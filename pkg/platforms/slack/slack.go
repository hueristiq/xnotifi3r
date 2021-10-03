package slack

import (
	"net/url"

	"github.com/containrrr/shoutrrr"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
)

type Platform struct {
	conf *configuration.Slack
}

func New(conf *configuration.Slack) (platform *Platform, err error) {
	platform = &Platform{
		conf: conf,
	}

	return
}

func (p *Platform) Send(message string) (err error) {
	url := &url.URL{
		Scheme: "slack",
		Path:   p.conf.Token + "@" + p.conf.Channel,
	}

	q := url.Query()

	if p.conf.Botname != "" {
		q.Set("botname", p.conf.Botname)
	}

	url.RawQuery = q.Encode()

	if err = shoutrrr.Send(url.String(), message); err != nil {
		return
	}

	return
}
