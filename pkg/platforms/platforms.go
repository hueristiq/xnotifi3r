package platforms

import (
	"strings"

	"github.com/hueristiq/hqnotifi3r/internal/configuration"
	"github.com/hueristiq/hqnotifi3r/pkg/platforms/slack"
	"github.com/hueristiq/hqnotifi3r/pkg/utils"
)

type Platform interface {
	Send(message string) error
}

type PlatformOptions struct {
	Slack *configuration.SlackConfiguration
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

	toUse := []string{}

	if strings.Contains(opts.Platform, ",") {
		toUse = append(toUse, strings.Split(opts.Platform, ",")...)
	} else if opts.Platform != "" {
		toUse = append(toUse, opts.Platform)
	}

	toIDS := []string{}

	if strings.Contains(opts.ID, ",") {
		toIDS = append(toIDS, strings.Split(opts.ID, ",")...)
	} else if opts.ID != "" {
		toIDS = append(toIDS, opts.ID)
	}

	if conf.PlatformsConfigurations.Slack != nil && (len(toUse) == 0 || utils.Contains(toUse, "slack")) {
		platform, err = slack.New(conf.PlatformsConfigurations.Slack, toIDS)
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
