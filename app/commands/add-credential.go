package commands

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CreateAddCredentialCommand() cli.Command {
	return cli.Command{
		Name:
		"add-credential",
		Aliases: []string{"a"},
		Usage:   "Add kubernetes credentials",
		Action: func(c *cli.Context) {
			AddCredential(c.String("name"),
				c.String("host"),
				c.String("username"),
				c.String("password"))
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "Kubernetes credential name",
			},
			cli.StringFlag{
				Name:  "host",
				Usage: "Kubernetes Host address",
			},
			cli.StringFlag{
				Name:  "username",
				Usage: "Kubernetes username",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Kubernetes password",
			},
		},
	}
}

func AddCredential(kubeName string, host string, username string, password string) {
	viper.Set(kubeName+":host", host)
	viper.Set(kubeName+":username", username)
	viper.Set(kubeName+":password", password)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saved your credentials")
	}
}
