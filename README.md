# Kubernetes Port Finder
kube-pf is a command line tool that find available ports on kubernetes. Written in Golang

## Getting Started

### Manual Installation

```bash
$ git clone https://github.com/trendyol/kube-port-finder

$ cd kube-port-finder

$ go build kube-pf.go
```

## Usage

```bash
$ ./kube-pf

NAME:
   kube-pf - Find available kubernetes node port number

USAGE:
   kube-pf [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Baris Ceviz @PeaceCwz

COMMANDS:
     add-credential, a  Add kubernetes credentials
     find, f            Find available port number
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Add Kubernetes Credential

```bash
$ ./kube-pf add-credential --name "your-kube-config-name" --host "https://127.0.0.1" --username "kube-user" --password "kube-password"

Saved your credentials
```

### Find Port Number

```bash
$ ./kube-pf find --name "your-kube-config-name" --range 1000,4000

Available 2292 port in your kubernetes
```

## Dependencies

* Viper [github.com/spf13/viper](github.com/spf13/viper)
* CLI [github.com/urfave/cli](github.com/urfave/cli)

## Contributing

* If you want to contribute to codes, create pull request
* If you find any bugs or error, create an issue

## License

This project is licensed under the MIT Lıcense