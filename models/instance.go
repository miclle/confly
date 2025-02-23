package models

type Instance struct {
	ID                string   `bson:"id"                 json:"id"`
	NamespaceName     string   `bson:"namespace_name"     json:"namespace_name"`
	ServiceName       string   `bson:"service_name"       json:"service_name"`
	ApplicationName   string   `bson:"application_name"   json:"application_name"`
	ConfigurationName string   `bson:"configuration_name" json:"configuration_name"`
	Environment       string   `bson:"environment"        json:"environment"`
	ProcessName       string   `bson:"process_name"       json:"process_name"`
	Hostname          string   `bson:"hostname"           json:"hostname"`
	IP                string   `bson:"ip"                 json:"ip"`
	Tags              []string `bson:"tags"               json:"tags"`
	VersionID         string   `bson:"version_id"         json:"version_id"`
	CreatedAt         int64    `bson:"created_at"         json:"created_at"`
	UpdatedAt         int64    `bson:"updated_at"         json:"updated_at"`
}
