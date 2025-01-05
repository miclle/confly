package models

type PublishType string

const (
	PublishTypeFullRelease PublishType = "FULL_RELEASE"
	PublishTypeGrayRelease PublishType = "GRAY_RELEASE"
)

type GrayRule struct {
	AllowHosts []string `bson:"allowHosts" json:"allowHosts"`
	AllowTags  []string `bson:"allowTags"  json:"allowTags"`
	AllowIPs   []string `bson:"allowIPs"   json:"allowIPs"`
}

type Publish struct {
	ID          string `bson:"id"          json:"id"`
	GroupID     string `bson:"groupID"     json:"groupID"`
	AppID       string `bson:"appID"       json:"appID"`
	ConfigSetID string `bson:"configSetID" json:"configSetID"`
	ClusterName string `bson:"clusterName" json:"clusterName"`

	Title       string      `bson:"title"                json:"title"`
	Description string      `bson:"description"          json:"description"`
	Type        PublishType `bson:"type"                 json:"type"`
	RollbackID  *string     `bson:"rollbackID,omitempty" json:"rollbackID,omitempty"`
	GrayRule    *GrayRule   `bson:"grayRule,omitempty"   json:"grayRule,omitempty"`

	ConfigFormat   ConfigFormat `bson:"configFormat"   json:"configFormat"`
	ConfigContent  string       `bson:"configContent"  json:"configContent"`
	ConfigChecksum string       `bson:"configChecksum" json:"configChecksum"`

	CreatedBy string  `bson:"createdBy"           json:"createdBy"`
	CreatedAt int64   `bson:"createdAt"           json:"createdAt"`
	DeletedBy *string `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
	DeletedAt *int64  `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
