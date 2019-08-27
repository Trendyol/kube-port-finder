package commands

import (
	"fmt"
	utils "github.com/trendyol/kube-port-finder/app"
	"github.com/trendyol/kube-port-finder/app/services"
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

	var randomPort int
	for {
		randomPort = utils.Random(minPort, maxPort)
		if utils.Contains(allPorts, randomPort) == false {
			break
		}
	}

	fmt.Printf("Available %d port in your kubernetes\n", randomPort)
}
