package params

import "github.com/miclle/confly/models"

type PublishConfigSet struct {
	ConfigSetID    string
	Type           models.PublishType
	Title          string
	Description    string
	GrayRule       *models.GrayRule
	ConfigChecksum string
	CreatedBy      string
}

type RevertConfigSet struct {
	GroupName     string
	AppName       string
	ClusterName   string
	ConfigSetName string
	ConfigSetID   string
	CreatedBy     string
}

type RollbackConfigSet struct {
	GroupName     string
	AppName       string
	ClusterName   string
	ConfigSetName string
	ConfigSetID   string
	ToPublishID   string
	CreatedBy     string
}

type GetPublishes struct {
	models.Pagination[*models.Publish]

	Q             string
	GroupID       string
	AppID         string
	ConfigSetID   string
	ConfigSetName string
	Type          models.PublishType
}

type GetPublish struct {
	PublishID string
}
