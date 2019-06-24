package models

import "time"

type KubeServiceResult struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				App      string `json:"app"`
				Chart    string `json:"chart"`
				Heritage string `json:"heritage"`
				Release  string `json:"release"`
			} `json:"labels"`
			Annotations struct {
				PrometheusIoHTTPProbe string `json:"prometheus.io/http_probe"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Ports []struct {
				Name     string `json:"name"`
				Protocol string `json:"protocol"`
				Port     int    `json:"port"`
				//TargetPort string `json:"targetPort"`
				NodePort int `json:"nodePort"`
			} `json:"ports"`
			Selector struct {
				App     string `json:"app"`
				Release string `json:"release"`
			} `json:"selector"`
			ClusterIP             string `json:"clusterIP"`
			Type                  string `json:"type"`
			SessionAffinity       string `json:"sessionAffinity"`
			ExternalTrafficPolicy string `json:"externalTrafficPolicy"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct {
			} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
}
