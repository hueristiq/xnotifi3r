package configuration

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Slack struct {
	Enabled    bool   `yaml:"enabled"`
	WebHookURL string `yaml:"webhook_url"`
}

type Platforms struct {
	Slack Slack `yaml:"slack"`
}

type Configuration struct {
	Version   string    `yaml:"version"`
	Platforms Platforms `yaml:"platforms"`
}

type Options struct {
	ExcludeSources  string
	UseSources      string
	SlackWebHookURL string
	Text            string

	YAMLConfig Configuration
}

const VERSION = "1.0.0"

func (options *Options) Parse() (err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	confPath := dir + "/.config/signotifi3r/conf.yaml"

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		configuration := Configuration{
			Version: VERSION,
			Platforms: Platforms{
				Slack: Slack{
					Enabled:    true,
					WebHookURL: "",
				},
			},
		}

		directory, _ := path.Split(confPath)

		if err := makeDirectory(directory); err != nil {
			return err
		}

		if err = configuration.MarshalWrite(confPath); err != nil {
			return err
		}

		options.YAMLConfig = configuration
	} else {
		configuration, err := UnmarshalRead(confPath)
		if err != nil {
			return err
		}

		if configuration.Version != VERSION {
			configuration.Version = VERSION

			if err := configuration.MarshalWrite(confPath); err != nil {
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
func (c *Configuration) MarshalWrite(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Indent the spaces too
	enc := yaml.NewEncoder(f)
	enc.SetIndent(4)
	err = enc.Encode(&c)
	f.Close()
	return err
}

// UnmarshalRead reads the unmarshalled config yaml file from disk
func UnmarshalRead(file string) (Configuration, error) {
	config := Configuration{}

	f, err := os.Open(file)
	if err != nil {
		return config, err
	}

	err = yaml.NewDecoder(f).Decode(&config)

	f.Close()

	return config, err
}
