package runner

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/containrrr/shoutrrr"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/platforms"
)

type Runner struct {
	options   *configuration.Options
	platforms *platforms.Client
}

func New(options *configuration.Options) (runner *Runner, err error) {
	// Discard all internal logs
	shoutrrr.SetLogger(log.New(ioutil.Discard, "", 0))

	prClient, err := platforms.New(&options.YAMLConfig, options)
	if err != nil {
		return nil, err
	}

	return &Runner{options: options, platforms: prClient}, nil
}

func (r *Runner) Run() (err error) {
	var inFile *os.File

	switch {
	case hasStdin():
		inFile = os.Stdin

	case r.options.Data != "":
		inFile, err = os.Open(r.options.Data)
		if err != nil {
			return
		}
	default:
		return errors.New("signotifi3r works with stdin or file using '-d' flag")
	}

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		msg := scanner.Text()

		r.sendMessage(msg)

	}

	return
}

func (r *Runner) sendMessage(msg string) error {
	if len(msg) > 0 {
		err := r.platforms.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) Close() {}
