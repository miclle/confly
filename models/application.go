package models

type Application struct {
	ID          string  `bson:"id"                  json:"id"`
	NamespaceID string  `bson:"namespaceID"         json:"namespaceID"`
	ServiceID   string  `bson:"serviceID"           json:"serviceID"`
	Name        string  `bson:"name"                json:"name"`
	Description string  `bson:"description"         json:"description"`
	CreatedBy   string  `bson:"createdBy"           json:"createdBy"`
	CreatedAt   int64   `bson:"createdAt"           json:"createdAt"`
	UpdatedBy   string  `bson:"updatedBy"           json:"updatedBy"`
	UpdatedAt   int64   `bson:"updatedAt"           json:"updatedAt"`
	DeletedBy   *string `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
	DeletedAt   *int64  `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`

	Namespace *Namespace `bson:"-" json:"namespace"`
	Service   *Service   `bson:"-" json:"service"`
}
