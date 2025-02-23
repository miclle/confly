package params

import "github.com/miclle/confly/models"

type UpsertInstance struct {
	NamespaceName     string
	ServiceName       string
	ApplicationName   string
	ConfigurationName string
	Environment       string
	Hostname          string
	Tags              []string
	IP                string
	VersionID         string
}

type GetInstances struct {
	models.Pagination[*models.Instance]

	NamespaceName     string
	ServiceName       string
	ApplicationName   string
	ConfigurationName string
	Environment       string
	Q                 string
}
