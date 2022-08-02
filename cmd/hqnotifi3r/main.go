package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hueristiq/hqnotifi3r/internal/configuration"
	"github.com/hueristiq/hqnotifi3r/internal/runner"
)

var (
	options configuration.Options
)

func init() {
	flag.StringVar(&options.Data, "d", "", "")
	flag.StringVar(&options.Data, "data", "", "")

	flag.StringVar(&options.Platform, "p", "", "")
	flag.StringVar(&options.Platform, "platform", "", "")

	flag.StringVar(&options.ID, "i", "", "")
	flag.StringVar(&options.ID, "id", "", "")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, configuration.BANNER)

		h := "USAGE:\n"
		h += "  hqnotifi3r [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "  -d, --data            file path to read data from\n"
		h += "  -p, --platform        platform to send notification to\n"
		h += "  -i, --id              id to send the notification to\n"
		h += "\n"

		fmt.Fprint(os.Stderr, h)
	}

	flag.Parse()
}

func main() {
	if err := options.Parse(); err != nil {
		log.Fatalln(err)
	}

	r, err := runner.New(&options)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			log.Println("\r- Ctrl+C pressed in Terminal")
			r.Close()
			os.Exit(0)
		}()
	}()

	if err = r.Run(); err != nil {
		log.Fatalln(err)
	}
}
