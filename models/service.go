package models

type Service struct {
	ID          string  `bson:"id"                   json:"id"`
	NamespaceID string  `bson:"namespace_id"         json:"namespace_id"`
	Name        string  `bson:"name"                 json:"name"`
	Description string  `bson:"description"          json:"description"`
	CreatedBy   string  `bson:"created_by"           json:"created_by"`
	CreatedAt   int64   `bson:"created_at"           json:"created_at"`
	UpdatedBy   string  `bson:"updated_by"           json:"updated_by"`
	UpdatedAt   int64   `bson:"updated_at"           json:"updated_at"`
	DeletedBy   *string `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	DeletedAt   *int64  `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`

	Namespace *Namespace `bson:"-" json:"namespace"`
}
