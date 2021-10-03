package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/logrusorgru/aurora/v3"
	"github.com/signedsecurity/signotifi3r/internal/configuration"
	"github.com/signedsecurity/signotifi3r/internal/runner"
)

var (
	options configuration.Options
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
	flag.StringVar(&options.Data, "d", "", "")
	flag.StringVar(&options.Data, "data", "", "")

	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  signotifi3r [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "  -d, --data        file path to read data from\n"

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

	// Setup close handler
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
