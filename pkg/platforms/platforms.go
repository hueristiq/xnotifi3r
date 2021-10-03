package platforms

import (
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/ansi"
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

	if conf.Platforms.Slack != nil {
		platform, err = slack.New(conf.Platforms.Slack)
		if err != nil {
			return
		}

		client.platforms = append(client.platforms, platform)
	}

	return
}

func (p *Client) Send(message string) (err error) {
	message = ansi.Strip(message)

	for _, platform := range p.platforms {
		if err = platform.Send(message); err != nil {
			return
		}
	}

	return
}
