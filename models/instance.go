package models

type Instance struct {
	ID                string   `bson:"id"                 json:"id"`
	NamespaceName     string   `bson:"namespace_name"     json:"namespace_name"`
	AppName           string   `bson:"app_name"           json:"app_name"`
	ConfigurationName string   `bson:"configuration_name" json:"configuration_name"`
	ClusterName       string   `bson:"cluster_name"       json:"cluster_name"`
	ServiceName       string   `bson:"service_name"       json:"service_name"`
	Hostname          string   `bson:"hostname"           json:"hostname"`
	Tags              []string `bson:"tags"               json:"tags"`
	IP                string   `bson:"ip"                 json:"ip"`
	PublishID         string   `bson:"publish_id"         json:"publish_id"`
	CreatedAt         int64    `bson:"created_at"         json:"created_at"`
	UpdatedAt         int64    `bson:"updated_at"         json:"updated_at"`
}
