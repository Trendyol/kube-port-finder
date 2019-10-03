package commands

import (
	"../services"
	"../utils"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

func CreateFindCommand() cli.Command {
	kubeService := new(services.KubernetesService)
	return cli.Command{
		Name:    "find",
		Aliases: []string{"f"},
		Usage:   "Find available port number",
		Action: func(c *cli.Context) {
			rangeString := strings.Split(c.String("range"), ",")
			minPort, err := strconv.Atoi(rangeString[0])

			if err != nil {
				fmt.Println("Invaild minimum port")
			}

			maxPort, err := strconv.Atoi(rangeString[1])

			if err != nil {
				fmt.Println("Invaild maximum port")
			}

			FindPorts(kubeService, c.String("name"), minPort, maxPort)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "Kubernetes credential name",
			},
			cli.StringFlag{
				Name:  "range",
				Usage: "Minimum and maximum port range",
			},
		},
	}
}

func FindPorts(service *services.KubernetesService, kubeName string, minPort int, maxPort int) {
	host := viper.GetString(kubeName + ":host")
	username := viper.GetString(kubeName + ":username")
	password := viper.GetString(kubeName + ":password")

	result, err := service.GetAllServices(host, username, password)
	if err != nil {
		fmt.Println(err)
	}

	var allPorts []int

	for _, app := range result {
		for _, port := range app.Ports {
			allPorts = append(allPorts, port)
		}
	}

	availablePort := -1

	for i := minPort; i <= maxPort; i++ {
		if utils.Contains(allPorts, i) == false {
			availablePort = i
			break
		}
	}

	fmt.Printf("Available %d port in your kubernetes\n", availablePort)
}
