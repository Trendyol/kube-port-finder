package services

import (
	b64 "encoding/base64"
	utils "github.com/trendyol/kube-port-finder/app"
	"github.com/trendyol/kube-port-finder/app/models"
)

type KubeService struct {
	AppName string
	Ports   []int
}

type KubernetesService struct {
}

func (service KubernetesService) GetAllServices(host string, username string, password string) ([]KubeService, error) {
	var result []KubeService
	auth := b64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	authHeader := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + auth,
	}
	var kubeServicesResponse = models.KubeServiceResult{}
	err := utils.SendRequest(host+"/api/v1/services/", authHeader, &kubeServicesResponse)

	if err != nil {
		return result, err
	}

	for _, item := range kubeServicesResponse.Items {
		var ports []int

		for _, port := range item.Spec.Ports {
			ports = append(ports, port.NodePort)
		}

		var kubeService = KubeService{
			AppName: item.Metadata.Name,
			Ports:   ports,
		}
		result = append(result, kubeService)
	}

	return result, err
}
