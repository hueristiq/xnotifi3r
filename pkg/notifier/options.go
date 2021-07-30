package notifier

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Version   string `yaml:"version"`
	Platforms struct {
		Slack struct {
			Use        bool   `yaml:"use"`
			WebHookURL string `yaml:"webhook_url"`
		} `yaml:"slack"`
	} `yaml:"platforms"`
}

type Options struct {
	ExcludeSources  string
	UseSources      string
	SlackWebHookURL string
	Text            string

	YAMLConfig Configuration
}

var version = "1.0.0"

func ParseOptions(options *Options) (*Options, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return options, err
	}

	confPath := dir + "/.config/signotifi3r/conf.yaml"

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		// first run
		configuration := Configuration{
			Version: version,
		}

		directory, _ := path.Split(confPath)

		err := makeDirectory(directory)
		if err != nil {
			return options, err
		}

		err = configuration.MarshalWrite(confPath)
		if err != nil {
			return options, err
		}

		options.YAMLConfig = configuration
	} else {
		// normal run
		configuration, err := UnmarshalRead(confPath)
		if err != nil {
			return options, err
		}

		if configuration.Version != version {
			configuration.Version = version

			err := configuration.MarshalWrite(confPath)
			if err != nil {
				return options, err
			}
		}

		options.YAMLConfig = configuration
	}

	return options, nil
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
