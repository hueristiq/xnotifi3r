package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/logrusorgru/aurora/v3"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/pkg/notifier"
)

type options struct {
	oneline bool
}

var (
	co options
	so configuration.Options
)

func banner() {
	fmt.Fprintln(os.Stderr, aurora.BrightBlue(`
     _                   _   _  __ _ _____      
 ___(_) __ _ _ __   ___ | |_(_)/ _(_)___ / _ __ 
/ __| |/ _`+"`"+` | '_ \ / _ \| __| | |_| | |_ \| '__|
\__ \ | (_| | | | | (_) | |_| |  _| |___) | |   
|___/_|\__, |_| |_|\___/ \__|_|_| |_|____/|_| v1.0.0
       |___/
`).Bold())
}

func init() {
	flag.BoolVar(&co.oneline, "l", false, "")

	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  signotifi3r [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "  -l        send message line by line (default: false)\n"

		fmt.Fprint(os.Stderr, h)
	}

	flag.Parse()
}

func main() {
	if err := so.Parse(); err != nil {
		log.Fatalln(err)
	}

	notifier, err := notifier.New(&so)
	if err != nil {
		log.Fatalln(err)
	}

	var lines string
	var message string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		message = line

		if co.oneline {
			notifier.SendNotification(message)
		} else {
			lines += line
			lines += "\n"
		}
	}

	if !co.oneline {
		message = lines

		notifier.SendNotification(message)
	}
}
