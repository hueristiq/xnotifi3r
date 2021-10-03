package configuration

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Slack struct {
	Token   string `yaml:"token"`
	Botname string `yaml:"botname"`
	Channel string `yaml:"channel"`
}

type Platforms struct {
	Slack *Slack `yaml:"slack"`
}

type Configuration struct {
	Version   string     `yaml:"version"`
	Platforms *Platforms `yaml:"platforms"`
}

type Options struct {
	Data string

	YAMLConfig *Configuration
}

const VERSION = "1.0.0"

func (options *Options) Parse() (err error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	confPath := dir + "/.config/signotifi3r/conf.yaml"

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		configuration := &Configuration{
			Version: VERSION,
			Platforms: &Platforms{
				Slack: &Slack{
					Token:   "",
					Botname: "",
					Channel: "",
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
func UnmarshalRead(file string) (configuration *Configuration, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&configuration)

	return
}
