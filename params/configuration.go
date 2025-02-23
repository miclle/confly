package params

import "github.com/miclle/confly/models"

type CreateConfiguration struct {
	ApplicationID string
	Name          string
	Description   string
	ConfigFormat  models.ConfigurationFormat
	CreatedBy     string
}

type GetConfigurations struct {
	models.Pagination[*models.Configuration]

	Q             string
	ApplicationID string
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
	ApplicationName   string
	Environment       string
	ConfigurationName string
	Hostname          string
	IP                string
	Tags              []string
	VersionID         string
}

type PublishConfiguration struct {
	ConfigurationID string
	Type            models.ReleaseType
	Title           string
	Description     string
	GrayRule        *models.ReleaseGrayRule
	ConfigChecksum  string
	CreatedBy       string
}

type RevertConfiguration struct {
	NamespaceName     string
	ApplicationName   string
	Environment       string
	ConfigurationName string
	ConfigurationID   string
	CreatedBy         string
}

type RollbackConfiguration struct {
	NamespaceName     string
	ApplicationName   string
	Environment       string
	ConfigurationName string
	ConfigurationID   string
	ToVersionID       string
	CreatedBy         string
}

type GetReleases struct {
	models.Pagination[*models.ConfigurationRelease]

	Q                 string
	NamespaceID       string
	ApplicationID     string
	ConfigurationID   string
	ConfigurationName string
	Type              models.ReleaseType
}

type GetRelease struct {
	VersionID string
}
