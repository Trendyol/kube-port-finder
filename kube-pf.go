package main

import (
	cmd "./app/commands"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	initConfig()
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "kube-pf"
	app.Usage = "Find available kubernetes node port number"
	app.Author = "Baris Ceviz @PeaceCwz"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		cmd.CreateAddCredentialCommand(),
		cmd.CreateFindCommand(),
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
