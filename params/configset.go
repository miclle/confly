package params

import "github.com/miclle/confly/models"

type CreateConfigSet struct {
	AppID        string
	Name         string
	Description  string
	ConfigFormat models.ConfigFormat
	CreatedBy    string
}

type GetConfigSets struct {
	models.Pagination[*models.ConfigSet]

	Q     string
	AppID string
}

type UpdateConfigSet struct {
	Description *string
	UpdatedBy   string
}

type UpdateConfigContent struct {
	ConfigContent  string
	ConfigChecksum string
	UpdatedBy      string
}
