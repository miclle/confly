package params

import "github.com/miclle/confly/models"

type PublishConfiguration struct {
	ConfigurationID string
	Type            models.PublishType
	Title           string
	Description     string
	GrayRule        *models.GrayRule
	ConfigChecksum  string
	CreatedBy       string
}

type RevertConfiguration struct {
	NamespaceName     string
	AppName           string
	ClusterName       string
	ConfigurationName string
	ConfigurationID   string
	CreatedBy         string
}

type RollbackConfiguration struct {
	NamespaceName     string
	AppName           string
	ClusterName       string
	ConfigurationName string
	ConfigurationID   string
	ToPublishID       string
	CreatedBy         string
}

type GetPublishes struct {
	models.Pagination[*models.Publish]

	Q                 string
	NamespaceID       string
	ApplicationID     string
	ConfigurationID   string
	ConfigurationName string
	Type              models.PublishType
}

type GetPublish struct {
	PublishID string
}
