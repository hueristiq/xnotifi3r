package platforms

import (
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/utils"
	"github.com/signedsecurity/signotifi3r/pkg/platforms/slack"
)

type Platform interface {
	Send(message string) error
}

type PlatformOptions struct {
	Slack *configuration.Slack
}

type Client struct {
	platforms []Platform
	conf      *configuration.Configuration
	opts      *configuration.Options
}

func New(conf *configuration.Configuration, opts *configuration.Options) (client *Client, err error) {
	var platform Platform

	client = &Client{
		conf: conf,
	}

	if conf.PlatformsConfigurations.Slack != nil && (len(conf.Platforms) == 0 || utils.Contains(conf.Platforms, "slack")) {
		platform, err = slack.New(conf.PlatformsConfigurations.Slack)
		if err != nil {
			return
		}

		client.platforms = append(client.platforms, platform)
	}

	return
}

func (p *Client) Send(message string) (err error) {
	message = utils.StripANSI(message)

	for _, platform := range p.platforms {
		if err = platform.Send(message); err != nil {
			return
		}
	}

	return
}
