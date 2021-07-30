# signotifi3r

![made with go](https://img.shields.io/badge/made%20with-Go-0040ff.svg) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/signedsecurity/signotifi3r.svg?style=flat&color=0040ff)](https://github.com/signedsecurity/signotifi3r/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/signedsecurity/signotifi3r.svg?style=flat&color=0040ff)](https://github.com/signedsecurity/signotifi3r/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/License-MIT-gray.svg?colorB=0040FF)](https://github.com/signedsecurity/signotifi3r/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@signedsecurity-0040ff.svg)](https://twitter.com/signedsecurity)

signotifi3r is a CLI helper utility to send notification via webhooks to Slack.

## Resources

* [Installation](#installation)
    * [From Binary](#from-binary)
    * [From Source](#from-source)
    * [From Github](#from-github)
* [Post Install Setup](#post-install-setup)
    * [Slack](#slack)
    * [Config File](#config-file)
* [Usage](#usage)

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/signedsecurity/signotifi3r/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

signotifi3r requires **go1.14+** to install successfully. Run the following command to get the repo

```bash
GO111MODULE=on go get github.com/signedsecurity/signotifi3r/cmd/signotifi3r
```

#### From Github

```bash
git clone https://github.com/signedsecurity/signotifi3r.git && \
cd signotifi3r/cmd/signotifi3r/ && \
go build . && \
mv signotifi3r /usr/local/bin/ && \
signotifi3r -h
```

## Post Insall Setup 

#### Slack

* __step 1:__ [Get yours Slack incoming webhook URL](https://slack.com/intl/en-id/help/articles/115005265063-Incoming-webhooks-for-Slack)
* __step 2:__ Add the webhook URL to the config file.

#### Config File

The default config file should be located in `$HOME/.config/signotifi3r/conf.yaml` and has the following contents:

```yaml
version: 1.0.0
platforms:
    slack:
        use: true
        webhook_url: "https://hooks.slack.com/services/xxxxxxxxxxx/xxxxxxxxxxx/xxxxxxxxxxx"
```

## Usage

To display help message for sigurls use the `-h` flag:

```bash
signotifi3r -h
```

```
     _                   _   _  __ _ _____      
 ___(_) __ _ _ __   ___ | |_(_)/ _(_)___ / _ __ 
/ __| |/ _` | '_ \ / _ \| __| | |_| | |_ \| '__|
\__ \ | (_| | | | | (_) | |_| |  _| |___) | |   
|___/_|\__, |_| |_|\___/ \__|_|_| |_|____/|_| v1.0.0
       |___/

USAGE:
  signotifi3r [OPTIONS]

OPTIONS:
  -l        send message line by line (default: false)
```

## Contribution

[Issues](https://github.com/signedsecurity/signotifi3r/issues) and [Pull Requests](https://github.com/signedsecurity/signotifi3r/pulls) are welcome!