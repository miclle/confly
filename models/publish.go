package models

type PublishType string

const (
	PublishTypeFullRelease PublishType = "FULL_RELEASE"
	PublishTypeGrayRelease PublishType = "GRAY_RELEASE"
)

type GrayRule struct {
	AllowHosts []string `bson:"allow_hosts" json:"allow_hosts"`
	AllowTags  []string `bson:"allow_tags"  json:"allow_tags"`
	AllowIPs   []string `bson:"allow_ips"   json:"allow_ips"`
}

type Publish struct {
	ID              string `bson:"id"               json:"id"`
	NamespaceID     string `bson:"namespace_id"     json:"namespace_id"`
	ApplicationID   string `bson:"application_id"   json:"application_id"`
	ConfigurationID string `bson:"configuration_id" json:"configuration_id"`
	ClusterName     string `bson:"cluster_name"     json:"cluster_name"`

	Title       string      `bson:"title"                 json:"title"`
	Description string      `bson:"description"           json:"description"`
	Type        PublishType `bson:"type"                  json:"type"`
	RollbackID  *string     `bson:"rollback_id,omitempty" json:"rollback_id,omitempty"`
	GrayRule    *GrayRule   `bson:"gray_rule,omitempty"   json:"gray_rule,omitempty"`

	ConfigFormat   ConfigFormat `bson:"config_format"   json:"config_format"`
	ConfigContent  string       `bson:"config_content"  json:"config_content"`
	ConfigChecksum string       `bson:"config_checksum" json:"config_checksum"`

	CreatedBy string `bson:"created_by" json:"created_by"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedBy string `bson:"updated_by" json:"updated_by"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}
