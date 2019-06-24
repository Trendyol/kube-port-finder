# Kubernetes Port Finder
Find kubernetes available port finder CLI tool

## Getting Started

### Installation (OSX) - Not Ready

```bash
$ brew install kube-pf
```

### Installation (Ubuntu) - Not Ready

```bash
$ apt-get update && apt-get install kube-pf
```

### Installation (Windows) - Not Ready

```
choco install kube-pf
```

### Manual Installation

```bash
$ go get -u https://github.com/trendyol/kube-port-finder

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
```

### Find Port Number

```bash
$ ./kube-pf find --name "your-kube-config-name" --range 1000,4000
```

## Contributing

* If you want to contribute to codes, create pull request
* If you find any bugs or error, create an issue

## License

This project is licensed under the MIT Lıcense