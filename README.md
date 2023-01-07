# Crog

CLI tool to setup scheduled tasks and call URLs based on the result, configured in yaml.

## Installation

### Debian/Ubuntu

1. Download the latest deb for your system from the [releases](https://github.com/henrywhitaker3/crog/releases) page.
2. Install the package: `sudo dpkg -i <name>.deb`
3. Add your checks to `/etc/crog/crog.yaml`
4. Start the service `sudo systemctl start crog.service`
5. Enable the service to start on boot: `sudo systemctl enable crog.service`

### Build from source

1. Clone the repo and navigate to the directory in your terminal
2. Run `CGO_ENABLED=0 go build -o crog`
3. Create a config file in `/etc/crog/crog.yaml` (see `build/crog.yaml` for empty default file)
4. Copy the service file in `build/crog.service` to `/etc/systemd/system/crog.service`
5. Run `sudo systemctl daemon-reload`
6. Start and enable the service

## Configuration

The config is defined in yaml, and the systemd service looks at `/etc/crog/crog.yaml` by default. Here is an example file with 1 check that pings `8.8.8.8` every minute, and mark the healtchecks.io check failed if it cannot reach it:

```yaml
version: 1
verbose: true # enables verbose log output
timezone: UTC

checks:
  - name: Ping google dns # The name of the check
    command: ping -W 1 -c 1 8.8.8.8 # The command to run
    code: 0 # The expected return code
    cron: "* * * * *" # The schedule the command will run on
    on:
      start: https://example.com/ping-google-dns/start
      success: https://example.com/ping-google-dns
      failure: https://example.com/ping-google-dns/fail
```

## Usage

The system service will run `work` as root, which will run the checks on a schedule. You can also run a check as a one-off by running `crog run` and choosing the check from the menu.
