package params

import "github.com/miclle/confly/models"

type CreateGroup struct {
	Name        string
	Description string
	CreatedBy   string
}

type GetGroups struct {
	models.Pagination[*models.Group]

	Q string
}

type GetGroup struct {
	Name string
	ID   string
}

type UpdateGroup struct {
	Description *string
	UpdatedBy   string
}
