package configuration

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/logrusorgru/aurora/v3"

	"gopkg.in/yaml.v3"
)

type SlackConfiguration struct {
	ID             string `yaml:"id"`
	SlackToken     string `yaml:"slack_token"`
	SlackBotname   string `yaml:"slack_botname"`
	SlackChannelID string `yaml:"slack_channel_id"`
}

type PlatformsConfigurations struct {
	Slack []*SlackConfiguration `yaml:"slack"`
}

type Configuration struct {
	Version                 string                   `yaml:"version"`
	PlatformsConfigurations *PlatformsConfigurations `yaml:"platforms"`
}

type Options struct {
	ID       string
	Data     string
	Platform string

	YAMLConfig Configuration
}

const (
	VERSION  string = "1.1.0"
	FILESDIR string = "hqnotifi3r"
)

var (
	BANNER string = aurora.Sprintf(
		aurora.BrightBlue(`
 _                       _   _  __ _ _____
| |__   __ _ _ __   ___ | |_(_)/ _(_)___ / _ __
| '_ \ / _`+"`"+` | '_ \ / _ \| __| | |_| | |_ \| '__|
| | | | (_| | | | | (_) | |_| |  _| |___) | |
|_| |_|\__, |_| |_|\___/ \__|_|_| |_|____/|_| %s
          |_|

`).Bold(),
		aurora.BrightRed("v"+VERSION).Bold(),
	)
	FILES string = func(folder string) string {
		dotConfig, err := os.UserConfigDir()
		if err != nil {
			log.Fatalln(err)
		}

		return filepath.Join(dotConfig, folder)
	}(FILESDIR)
)

func (options *Options) Parse() (err error) {
	configuration := Configuration{}
	confYAMLFile := filepath.Join(FILES, "conf.yaml")

	_, err = os.Stat(confYAMLFile)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil

			configuration = Configuration{
				PlatformsConfigurations: &PlatformsConfigurations{
					Slack: []*SlackConfiguration{
						{
							ID:             "slack",
							SlackToken:     "",
							SlackBotname:   "",
							SlackChannelID: "",
						},
					},
				},
				Version: VERSION,
			}

			directory, _ := path.Split(confYAMLFile)

			if err := makeDirectory(directory); err != nil {
				return err
			}

			if err = configuration.MarshalWrite(confYAMLFile); err != nil {
				return err
			}

			options.YAMLConfig = configuration
		} else {
			return
		}
	} else {
		configuration, err := UnmarshalRead(confYAMLFile)
		if err != nil {
			return err
		}

		if configuration.Version != VERSION {
			configuration.Version = VERSION

			if err := configuration.MarshalWrite(confYAMLFile); err != nil {
				return err
			}
		}

		options.YAMLConfig = configuration
	}

	return
}

func makeDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if directory != "" {
			err = os.MkdirAll(directory, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// MarshalWrite writes the marshaled yaml config to disk
func (c *Configuration) MarshalWrite(file string) (err error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}

	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(4)
	err = enc.Encode(&c)

	return
}

// UnmarshalRead reads the unmarshalled config yaml file from disk
func UnmarshalRead(file string) (configuration Configuration, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&configuration)

	return
}
