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

type ConfigSet struct {
	ID             string       `bson:"id"                  json:"id"`
	AppID          string       `bson:"appID"               json:"appID"`
	Name           string       `bson:"name"                json:"name"`
	Description    string       `bson:"description"         json:"description"`
	ClusterName    string       `bson:"clusterName"         json:"clusterName"`
	ConfigFormat   ConfigFormat `bson:"configFormat"        json:"configFormat"`
	ConfigContent  string       `bson:"configContent"       json:"configContent"`
	ConfigChecksum string       `bson:"configChecksum"      json:"configChecksum"`
	CreatedBy      string       `bson:"createdBy"           json:"createdBy"`
	CreatedAt      int64        `bson:"createdAt"           json:"createdAt"`
	UpdatedBy      string       `bson:"updatedBy"           json:"updatedBy"`
	UpdatedAt      int64        `bson:"updatedAt"           json:"updatedAt"`
	DeletedBy      *string      `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
	DeletedAt      *int64       `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

type Configuration struct {
	GroupName     string       `json:"groupName"`
	AppName       string       `json:"appName"`
	ConfigSetName string       `json:"configSetName"`
	ClusterName   string       `json:"clusterName"`
	PublishID     string       `json:"publishID"`
	ConfigFormat  ConfigFormat `json:"configFormat"`
	ConfigContent string       `json:"configContent"`
}
