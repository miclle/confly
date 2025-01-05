package models

type Pagination[T any] struct {
	Items []T `bson:"items" json:"items"`
	Total int `bson:"total" json:"total"`
	Size  int `bson:"size"  json:"size"`
	Page  int `bson:"page"  json:"page"`
}
