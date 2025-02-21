package models

type Instance struct {
	ID                string   `bson:"id"                json:"id"`
	NamespaceName     string   `bson:"namespaceID"       json:"namespaceID"`
	AppName           string   `bson:"appName"           json:"appName"`
	ConfigurationName string   `bson:"configurationName" json:"configurationName"`
	ClusterName       string   `bson:"clusterName"       json:"clusterName"`
	ServiceName       string   `bson:"serviceName"       json:"serviceName"`
	Hostname          string   `bson:"hostname"          json:"hostname"`
	Tags              []string `bson:"tags"              json:"tags"`
	IP                string   `bson:"ip"                json:"ip"`
	PublishID         string   `bson:"publishID"         json:"publishID"`
	CreatedAt         int64    `bson:"createdAt"         json:"createdAt"`
	UpdatedAt         int64    `bson:"updatedAt"         json:"updatedAt"`
}
