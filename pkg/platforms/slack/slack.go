package slack

import (
	"net/url"

	"github.com/containrrr/shoutrrr"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/utils"
)

type Platform struct {
	Confs []*configuration.SlackConfiguration
}

func New(conf []*configuration.SlackConfiguration, toIDS []string) (platform *Platform, err error) {
	platform = &Platform{}

	for _, o := range conf {
		if len(toIDS) == 0 || utils.Contains(toIDS, o.ID) {
			platform.Confs = append(platform.Confs, o)
		}
	}

	return
}

func (platform *Platform) Send(message string) (err error) {
	for _, pr := range platform.Confs {
		url := &url.URL{
			Scheme: "slack",
			Path:   pr.SlackToken + "@" + pr.SlackChannelID,
		}

		q := url.Query()

		if pr.SlackBotname != "" {
			q.Set("botname", pr.SlackBotname)
		}

		url.RawQuery = q.Encode()

		if err = shoutrrr.Send(url.String(), message); err != nil {
			continue
		}
	}

	return
}
