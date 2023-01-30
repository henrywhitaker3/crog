# Crog

CLI tool to setup scheduled tasks and call URLs based on the result, configured in yaml.

## Installation

### Debian/Ubuntu

1. Download the latest deb for your system from the [releases](https://github.com/henrywhitaker3/crog/releases) page.
2. Install the package: `sudo dpkg -i <name>.deb`
3. Add your actions to `/etc/crog/crog.yaml`
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

The config is defined in yaml, and the systemd service looks at `/etc/crog/crog.yaml` by default. Here is an example file with 1 check that pings `8.8.8.8` every minute, and mark a healtchecks.io check failed if it cannot reach it:

```yaml
version: 1

actions:
  - name: Ping google dns
    command: ping -W 1 -c 1 8.8.8.8
    when:
      success: https://example.com/ping-google-dns
```

### Values

| Field | Description | Type | Default | Required |
| --- | --- | --- | --- | --- |
| version | The config format version | int |  | yes |
| timezone | The timezone the cron scheduler uses | string | UTC | no |
| verbose | Whether to turn on verbose logging | bool | false | no |
| actions | An array of actions to run | [] | | yes |
| actions.*.name | The name of the action | string | | yes |
| actions.*.command | The command to run | string | | yes |
| actions.*.code | The desired exit code returned by the command | int | 0 | no |
| actions.*.cron | The schedule that the action will be run on | string | * * * * * | no |
| actions.*.when | An object conatining the URLs to call after the actions are run | {} | | no |
| actions.*.when.start | The URL to call before the action gets run | string | | no |
| actions.*.when.success | The URL to call when the action is successful | string | | no |
| actions.*.when.failure | The URL to call when the action fails | string | | no |
| remotes | An array of remote crog servers | [] | | no |
| remotes.*.name | The name of the remote crog server | string | | yes |
| remotes.*.url | The url of the remote crog server | string | | yes |

## Usage

The system service will run `work` as root, which will run the checks on a schedule. You can also run a check as a one-off by running `crog run` and choosing the check from the menu.

### Remote server

The crog service will also run a grpc server, which allows you to run crog actions on remote servers. By default, this is turned off. To enable the remote server, add the following to your crog config on the server:

```yaml
...
server:
    enabled: true
    listen: :9399 # You can chose any port here. Add an IP before the colon to only listen on a specific IP address
```

To run commands on remote servers, install crog where you want to run commands from and add these remote servers to your config file:

```yaml
...
remotes:
    - name: Example server
      url: 10.10.10.10:9399
```

Be careful when not setting `server.listen` to a specific IP, as there is currently no authentication and anyone could run these commands if your server has a public IP.
