This is a small application designed for triggering a IFTTT webhook whenever a keyword is found within a website's content.

## :wrench: &nbsp; Installation

The binary is available for Linux, macOS and Windows from the [GitHub Releases page](https://github.com/krakowski/webwatcher/releases).

## :rocket: &nbsp; Quick Start

The tool requires a configuration file which can be placed under `./config.yaml`. Alternatively, another path can be
specified using the `--config` program argument. The following configuration options must be specified.

```yaml
# The website to check.
website: "<FULL_WEBISTE_URL>"

# A string that must be present within the response body.
check: "Search Results"

# A list of keywords that will be matched against the response body.
keywords: ["Sale", "Buy"]

# The interval at which the website should be checked.
interval: "15s"

# The IFTTT webhook configuration. See https://ifttt.com/maker_webhooks under 
# "Documentation" to find out your personal key.
trigger:
  key: "<YOUR_IFTTT_WEBHOOK_KEY>"
  event: "<YOUR_WEBHOOK_EVENT_NAME>"
```

The tool can be started by executing the binary.

```console
./webwatcher
```

## :scroll: &nbsp; License

This project is licensed under the GNU GPLv3 License - see the [LICENSE](LICENSE) file for details.