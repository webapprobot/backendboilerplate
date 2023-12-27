# Backend-starter

Boilerplate code for building backend applications in go using [Webapp bot](https://github.com/webapprobot/webappbot)

~~Replace everything below. Make relevant for the current project~~
## Table of Contents

## Design
Refer to [Component Diagram](https://github.com/backendboilerplate/architecture/blob/master/APIcomponent.md)

## Requirements
1. postgres
2. An ubuntu server
3. A configured mail server 

## Development
```bash
git clone git@github.com:backendboilerplate/backend.git

##
cd backend
sudo sudo cp -r debData/etc /
sudo chown user:user /etc/backendboilerplate
```
### Configurations
Edit the files `/etc/backendboilerplate/backendboilerplateDev.yaml` with db and mail settings. The configuration for the `dev` environment is different from that for the production environments so that settings are not overwritten during installation.

Start the application
```bash
go run . run
```

## Building
You can build your application using `./build.sh`. This creates the executable in `./builds` directory.
*** check

## Installation
Download the latest release. [*Update link later*]()
```bash
wget download-link
```

Install
```bash
dpkg -i backendboilerplate
```

Once installed, install service using:
```bash
# dev env
sudo backendboilerplate init
# prod env
sudo backendboilerplate init --prod
```

## Configuration
The configuration file is found in `/etc/backendboilerplate/config.yaml`

## Usage
**In development**

Stop the service, which will be using the port, then run the command to run the program:

```bash
sudo systemctl stop backendboilerplateDev
# in your directory
go run . run
```


Starting, Stopping, Restarting service:
```bash
# dev env
sudo systemctl stop backendboilerplateDev
sudo systemctl start backendboilerplateDev
sudo systemctl restart backendboilerplateDev

# prod env
sudo systemctl stop backendboilerplate
sudo systemctl start backendboilerplate
sudo systemctl restart backendboilerplate
```
Or:

```bash
# dev env
sudo backendboilerplate start
sudo backendboilerplate stop
sudo backendboilerplate restart

# prod env
sudo backendboilerplate start --prod
sudo backendboilerplate stop --prod
sudo backendboilerplate restart --prod
```

Check the logs using

For `dev` Env
```bash
journalctl -f -u journalctl -f -u backendboilerplateDev
```

For `prod` Env
```bash
journalctl -f -u backendboilerplate
```

## Building
Check makefile for output dir.

```bash
make build
```

## Creating debs
Check makefile for output dir.

```bash
make debs
```

## Creating a Release
```bash
git tag v1.1.0
git push origin v1.1.0

```


## Installation

Download the relevant file for your system from [releases](https://github.com/webappbot/backendboilerplate/releases/tag/v1.0). Instructions may not work for private repos.

```
curl --silent "https://api.github.com/repos/backendboilerplate/API/releases/latest"|   grep "browser_download_url.*amd64.deb" | head -n 1 | cut -d : -f 2,3 | tr -d \"  | xargs wget -O tmp.deb && sudo dpkg -i tmp.deb
```


## Usage
1. Go to `/swagger/index.html`
2. There are a few details missing in swagger which you can find in [notes](notes.md)
