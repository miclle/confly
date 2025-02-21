package params

import "github.com/miclle/confly/models"

type UpsertInstance struct {
	NamespaceName     string
	ServiceName       string
	AppName           string
	ConfigurationName string
	ClusterName       string
	Hostname          string
	Tags              []string
	IP                string
	PublishID         string
}

type GetInstances struct {
	models.Pagination[*models.Instance]

	NamespaceName     string
	ServiceName       string
	AppName           string
	ConfigurationName string
	ClusterName       string
	Q                 string
}
