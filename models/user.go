package models

type User struct {
	ID        string  `bson:"id"                   json:"id"`
	Username  string  `bson:"username"             json:"username"`
	Name      string  `bson:"name"                 json:"name"`
	Email     string  `bson:"email"                json:"email"`
	CreatedBy string  `bson:"created_by"           json:"created_by"`
	CreatedAt int64   `bson:"created_at"           json:"created_at"`
	UpdatedBy string  `bson:"updated_by"           json:"updated_by"`
	UpdatedAt int64   `bson:"updated_at"           json:"updated_at"`
	DeletedBy *string `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	DeletedAt *int64  `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
