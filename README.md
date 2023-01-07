# Go-Healthcheck

CLI tool to configure healthchecks that interact with [healthcheck.io](https://healthchecks.io) configured in yaml.

## Installation

### Debian/Ubuntu

1. Download the latest deb for your system from the [releases](https://github.com/henrywhitaker3/go-healthcheck/releases) page.
2. Install the package: `sudo dpkg -i <name>.deb`
3. Add your checks to `/etc/healthchecks/healthchecks.yaml`
4. Start the service `sudo systemctl start healthcheck.service`
5. Enable the service to start on boot: `sudo systemctl enable healthcheck.service`

### Build from source

1. Clone the repo and navigate to the directory in your terminal
2. Run `CGO_ENABLED=0 go build -o healthcheck`
3. Create a config file in `/etc/healthcheck/healthcheck.yaml` (see `build/healthcheck.yaml` for empty default file)
4. Copy the service file in `build/healthcheck.service` to `/etc/systemd/system/healthcheck.service`
5. Run `sudo systemctl daemon-reload`
6. Start and enable the service

## Configuration

The config is defined in yaml, and the systemd service looks at `/etc/healthcheck/healthcheck.yaml` by default. Here is an example file with 1 check that pings `8.8.8.8` every minute, and mark the healtchecks.io check failed if it cannot reach it:

```yaml
version: 1
verbose: true # enables verbose log output
timezone: UTC

checks:
  - name: Ping google dns # The name of the check
    command: ping -W 1 -c 1 8.8.8.8 # The command to run
    code: 0 # The expected return code
    cron: "* * * * *" # The schedule the command will run on
    id: bc9e7f5f-dbba-40c5-82da-590cc70959e8 # The healthchecks.io check id
```

## Usage

The system service will run `work` as root, which will run the checks on a schedule. You can also run a check as a one-off by running `healthcheck run` and choosing the check from the menu.
