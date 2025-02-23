package models

type ConfigurationFormat string

var (
	ConfigFormatTXT        ConfigurationFormat = "txt"
	ConfigFormatJSON       ConfigurationFormat = "json"
	ConfigFormatYAML       ConfigurationFormat = "yaml"
	ConfigFormatXML        ConfigurationFormat = "xml"
	ConfigFormatProperties ConfigurationFormat = "properties"
	ConfigFormatTOML       ConfigurationFormat = "toml"
)

func (format ConfigurationFormat) String() string {
	return string(format)
}

type Configuration struct {
	ID string `bson:"id" json:"id"`

	NamespaceID     string `bson:"namespace_id"     json:"namespace_id"`
	NamespaceName   string `bson:"namespace_name"   json:"namespace_name"`
	ServiceID       string `bson:"service_id"       json:"service_id"`
	ServiceName     string `bson:"service_name"     json:"service_name"`
	ApplicationID   string `bson:"application_id"   json:"application_id"`
	ApplicationName string `bson:"application_name" json:"application_name"`

	Name        string              `bson:"name"        json:"name"`
	Description string              `bson:"description" json:"description"`
	Environment string              `bson:"environment" json:"environment"`
	Format      ConfigurationFormat `bson:"format"      json:"format"`
	Content     string              `bson:"content"     json:"content"`
	Checksum    string              `bson:"checksum"    json:"checksum"`

	CreatedBy string  `bson:"created_by"           json:"created_by"`
	CreatedAt int64   `bson:"created_at"           json:"created_at"`
	UpdatedBy string  `bson:"updated_by"           json:"updated_by"`
	UpdatedAt int64   `bson:"updated_at"           json:"updated_at"`
	DeletedBy *string `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	DeletedAt *int64  `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type ConfigurationHistory struct {
	ID              string              `bson:"id"               json:"id"`
	ConfigurationID string              `bson:"configuration_id" json:"configuration_id"`
	Format          ConfigurationFormat `bson:"format"           json:"format"`
	Content         string              `bson:"content"          json:"content"`
	Checksum        string              `bson:"checksum"         json:"checksum"`
	CreatedBy       string              `bson:"created_by"       json:"created_by"`
	CreatedAt       int64               `bson:"created_at"       json:"created_at"`
}

type ReleaseType string

const (
	ReleaseTypeFull        ReleaseType = "full"
	ReleaseTypeGrayRelease ReleaseType = "gray"
)

type ReleaseGrayRule struct {
	AllowHosts []string `bson:"allow_hosts" json:"allow_hosts"`
	AllowTags  []string `bson:"allow_tags"  json:"allow_tags"`
	AllowIPs   []string `bson:"allow_ips"   json:"allow_ips"`
}

type ConfigurationRelease struct {
	ID              string `bson:"id"               json:"id"`
	NamespaceID     string `bson:"namespace_id"     json:"namespace_id"`
	ApplicationID   string `bson:"application_id"   json:"application_id"`
	ConfigurationID string `bson:"configuration_id" json:"configuration_id"`
	Environment     string `bson:"environment"      json:"environment"`

	RollbackID  *string          `bson:"rollback_id,omitempty" json:"rollback_id,omitempty"`
	Title       string           `bson:"title"                 json:"title"`
	Description string           `bson:"description"           json:"description"`
	Type        ReleaseType      `bson:"type"                  json:"type"`
	GrayRule    *ReleaseGrayRule `bson:"gray_rule,omitempty"   json:"gray_rule,omitempty"`

	Format   ConfigurationFormat `bson:"format"   json:"format"`
	Content  string              `bson:"content"  json:"content"`
	Checksum string              `bson:"checksum" json:"checksum"`

	CreatedBy string `bson:"created_by" json:"created_by"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedBy string `bson:"updated_by" json:"updated_by"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}
