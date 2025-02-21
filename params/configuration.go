package params

import "github.com/miclle/confly/models"

type CreateConfiguration struct {
	AppID        string
	Name         string
	Description  string
	ConfigFormat models.ConfigFormat
	CreatedBy    string
}

type GetConfigurations struct {
	models.Pagination[*models.Configuration]

	Q     string
	AppID string
}

type UpdateConfiguration struct {
	Description *string
	UpdatedBy   string
}

type UpdateConfigurationContent struct {
	Content   string
	Checksum  string
	UpdatedBy string
}

type GetConfiguration struct {
	NamespaceName     string
	ServiceName       string
	AppName           string
	ClusterName       string
	ConfigurationName string
	Hostname          string
	IP                string
	Tags              []string
	PublishID         string
}
