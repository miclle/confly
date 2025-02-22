package models

type ConfigFormat string

var (
	ConfigFormatTXT        ConfigFormat = "txt"
	ConfigFormatJSON       ConfigFormat = "json"
	ConfigFormatYAML       ConfigFormat = "yaml"
	ConfigFormatXML        ConfigFormat = "xml"
	ConfigFormatProperties ConfigFormat = "properties"
	ConfigFormatTOML       ConfigFormat = "toml"
)

func (format ConfigFormat) String() string {
	return string(format)
}

type Configuration struct {
	ID            string       `bson:"id"                   json:"id"`
	NamespaceName string       `bson:"namespace_name"       json:"namespace_name"`
	ApplicationID string       `bson:"application_id"       json:"application_id"`
	Name          string       `bson:"name"                 json:"name"`
	Description   string       `bson:"description"          json:"description"`
	ClusterName   string       `bson:"cluster_name"         json:"cluster_name"`
	Format        ConfigFormat `bson:"format"               json:"format"`
	Content       string       `bson:"content"              json:"content"`
	Checksum      string       `bson:"checksum"             json:"checksum"`
	CreatedBy     string       `bson:"created_by"           json:"created_by"`
	CreatedAt     int64        `bson:"created_at"           json:"created_at"`
	UpdatedBy     string       `bson:"updated_by"           json:"updated_by"`
	UpdatedAt     int64        `bson:"updated_at"           json:"updated_at"`
	DeletedBy     *string      `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	DeletedAt     *int64       `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type ConfigurationVersion struct {
	ID              string       `bson:"id"               json:"id"`
	ConfigurationID string       `bson:"configuration_id" json:"configuration_id"`
	Format          ConfigFormat `bson:"format"           json:"format"`
	Content         string       `bson:"content"          json:"content"`
	Checksum        string       `bson:"checksum"         json:"checksum"`
	CreatedBy       string       `bson:"created_by"       json:"created_by"`
	CreatedAt       int64        `bson:"created_at"       json:"created_at"`
}
