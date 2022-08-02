# hqnotifi3r

![made with go](https://img.shields.io/badge/made%20with-Go-0040ff.svg) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hqnotifi3r.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqnotifi3r/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hqnotifi3r.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqnotifi3r/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/License-MIT-gray.svg?colorB=0040FF)](https://github.com/hueristiq/hqnotifi3r/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@hueristiq-0040ff.svg)](https://twitter.com/hueristiq)

hqnotifi3r is a helper utility to send messages from CLI to Slack.

## Resources

* [Installation](#installation)
	* [From Binary](#from-binary)
	* [From Source](#from-source)
	* [From Github](#from-github)
* [Post Install Setup](#post-install-setup)
	* [Config File](#config-file)
* [Usage](#usage)

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/hueristiq/hqnotifi3r/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

hqnotifi3r requires **go1.17+** to install successfully. Run the following command to get the repo

```bash
go install github.com/hueristiq/hqnotifi3r/cmd/hqnotifi3r@latest
```

#### From Github

```bash
git clone https://github.com/hueristiq/hqnotifi3r.git && \
cd hqnotifi3r/cmd/hqnotifi3r/ && \
go build . && \
mv hqnotifi3r /usr/local/bin/ && \
hqnotifi3r -h
```

## Post Insall Setup 

#### Config File

The default config file should be located in `$HOME/.config/hqnotifi3r/conf.yaml` and has the following contents:

```yaml
version: 1.0.0
platforms:
    slack:
        -
            id: "slack"
            slack_token: "xoxb-123456789012-1234567890123-4mt0t4l1YL3g1T5L4cK70k3N"
            slack_botname: "hqnotifi3r"
            slack_channel_id: "C039ZSYCYKT"
        -
            id: "targets"
            slack_token: "xoxb-123456789012-1234567890123-4mt0t4l1YL3g1T5L4cK70k3N"
            slack_botname: "hqnotifi3r"
            slack_channel_id: "C03L9RQRK4P"
```

## Usage

To display help message for sigurls use the `-h` flag:

```bash
hqnotifi3r -h
```

```
 _                       _   _  __ _ _____
| |__   __ _ _ __   ___ | |_(_)/ _(_)___ / _ __
| '_ \ / _` | '_ \ / _ \| __| | |_| | |_ \| '__|
| | | | (_| | | | | (_) | |_| |  _| |___) | |
|_| |_|\__, |_| |_|\___/ \__|_|_| |_|____/|_| v1.0.0
          |_|

USAGE:
  hqnotifi3r [OPTIONS]

OPTIONS:
  -d, --data            file path to read data from
  -p, --platform        platform to send notification to
  -i, --id              id to send the notification to
```

## Contribution

[Issues](https://github.com/hueristiq/hqnotifi3r/issues) and [Pull Requests](https://github.com/hueristiq/hqnotifi3r/pulls) are welcome!