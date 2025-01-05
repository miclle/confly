package params

import "github.com/miclle/confly/models"

type UpsertInstance struct {
	GroupName     string
	AppName       string
	ConfigSetName string
	ClusterName   string
	ServiceName   string
	Hostname      string
	Tags          []string
	IP            string
	PublishID     string
}

type GetInstances struct {
	models.Pagination[*models.Instance]

	Q             string
	GroupName     string
	AppName       string
	ConfigSetName string
	ClusterName   string
}
